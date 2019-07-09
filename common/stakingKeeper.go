package common

import (
	"encoding/hex"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/types"
)

// ----------------- Staking Related Keepers

// GetValidatorKey drafts the validator key for addresses
func GetValidatorKey(address []byte) []byte {
	return append(ValidatorsKey, address...)
}

// GetValidatorMapKey returns validator map
func GetValidatorMapKey(address []byte) []byte {
	return append(ValidatorMapKey, address...)
}

// AddValidator adds validator indexed with address
func (k *Keeper) AddValidator(ctx sdk.Context, validator types.Validator) error {
	// TODO uncomment
	//if ok:=validator.ValidateBasic(); !ok{
	//	// return error
	//}

	store := ctx.KVStore(k.StakingKey)

	bz, err := types.MarshallValidator(k.cdc, validator)
	if err != nil {
		return err
	}

	// store validator with address prefixed with validator key as index
	store.Set(GetValidatorKey(validator.Signer.Bytes()), bz)
	StakingLogger.Debug("Validator stored", "key", hex.EncodeToString(GetValidatorKey(validator.Signer.Bytes())), "validator", validator.String())

	// add validator to validator ID => SignerAddress map
	k.SetValidatorIDToSignerAddr(ctx, validator.ID, validator.Signer)
	return nil
}

// GetValidatorInfo returns validator
func (k *Keeper) GetValidatorInfo(ctx sdk.Context, address []byte) (validator types.Validator, err error) {
	store := ctx.KVStore(k.StakingKey)

	// check if validator exists
	key := GetValidatorKey(address)
	if !store.Has(key) {
		return validator, errors.New("Validator not found")
	}

	// unmarshall validator and return
	validator, err = types.UnmarshallValidator(k.cdc, store.Get(key))
	if err != nil {
		return validator, err
	}

	// return true if validator
	return validator, nil
}

// GetCurrentValidators returns all validators who are in validator set
func (k *Keeper) GetCurrentValidators(ctx sdk.Context) (validators []types.Validator) {
	// get ACK count
	ackCount := k.GetACKCount(ctx)

	// iterate through validator list
	k.IterateValidatorsAndApplyFn(ctx, func(validator types.Validator) error {
		// check if validator is valid for current epoch
		if validator.IsCurrentValidator(ackCount) {
			// append if validator is current valdiator
			validators = append(validators, validator)
		}
		return nil
	})

	return
}

// GetAllValidators returns all validators
func (k *Keeper) GetAllValidators(ctx sdk.Context) (validators []*types.Validator) {
	// iterate through validators and create validator update array
	k.IterateValidatorsAndApplyFn(ctx, func(validator types.Validator) error {
		// append to list of validatorUpdates
		validators = append(validators, &validator)
		return nil
	})

	return
}

// IterateValidatorsAndApplyFn interate validators and apply the given function.
func (k *Keeper) IterateValidatorsAndApplyFn(ctx sdk.Context, f func(validator types.Validator) error) {
	store := ctx.KVStore(k.StakingKey)

	// get validator iterator
	iterator := sdk.KVStorePrefixIterator(store, ValidatorsKey)
	defer iterator.Close()

	// loop through validators to get valid validators
	for ; iterator.Valid(); iterator.Next() {
		// unmarshall validator
		validator, _ := types.UnmarshallValidator(k.cdc, iterator.Value())
		// call function and return if required
		if err := f(validator); err != nil {
			return
		}
	}
}

// AddDeactivationEpoch adds deactivation epoch
func (k *Keeper) AddDeactivationEpoch(ctx sdk.Context, validator types.Validator, updatedVal types.Validator) error {
	// check if validator has unstaked
	if updatedVal.EndEpoch != 0 {
		validator.EndEpoch = updatedVal.EndEpoch
		// update validator in store
		k.AddValidator(ctx, validator)
		return nil
	}

	return errors.New("Deactivation period not set")
}

// UpdateSigner updates validator with signer and pubkey + validator => signer map
func (k *Keeper) UpdateSigner(ctx sdk.Context, newSigner common.Address, newPubkey types.PubKey, prevSigner common.Address) error {
	// get old validator from state and make power 0
	validator, err := k.GetValidatorInfo(ctx, prevSigner.Bytes())
	if err != nil {
		StakingLogger.Error("Unable to fetch valiator from store")
		return err
	}

	// copy power to reassign below
	validatorPower := validator.Power
	validator.Power = 0
	// update validator
	k.AddValidator(ctx, validator)

	//update signer in prev Validator
	validator.Signer = newSigner
	validator.PubKey = newPubkey
	validator.Power = validatorPower

	// add updated validator to store with new key
	k.AddValidator(ctx, validator)

	return nil
}

// UpdateValidatorSetInStore adds validator set to store
func (k *Keeper) UpdateValidatorSetInStore(ctx sdk.Context, newValidatorSet types.ValidatorSet) error {
	// TODO check if we may have to delay this by 1 height to sync with tendermint validator updates
	store := ctx.KVStore(k.StakingKey)

	// marshall validator set
	bz, err := k.cdc.MarshalBinaryBare(newValidatorSet)
	if err != nil {
		return err
	}

	// set validator set with CurrentValidatorSetKey as key in store
	store.Set(CurrentValidatorSetKey, bz)
	return nil
}

// GetValidatorSet returns current Validator Set from store
func (k *Keeper) GetValidatorSet(ctx sdk.Context) (validatorSet types.ValidatorSet) {
	store := ctx.KVStore(k.StakingKey)
	// get current validator set from store
	bz := store.Get(CurrentValidatorSetKey)
	// unmarhsall
	k.cdc.UnmarshalBinaryBare(bz, &validatorSet)

	// return validator set
	return validatorSet
}

// IncreamentAccum increments accum for validator set by n times and replace validator set in store
func (k *Keeper) IncreamentAccum(ctx sdk.Context, times int) {
	// get validator set
	validatorSet := k.GetValidatorSet(ctx)

	// increment accum
	validatorSet.IncrementAccum(times)

	// replace
	k.UpdateValidatorSetInStore(ctx, validatorSet)
}

// GetNextProposer returns next proposer
func (k *Keeper) GetNextProposer(ctx sdk.Context) *types.Validator {
	// get validator set
	validatorSet := k.GetValidatorSet(ctx)

	// Increment accum in copy
	copiedValidatorSet := validatorSet.CopyIncrementAccum(1)

	// get signer address for next signer
	return copiedValidatorSet.GetProposer()
}

// GetCurrentProposer returns current proposer
func (k *Keeper) GetCurrentProposer(ctx sdk.Context) *types.Validator {
	// get validator set
	validatorSet := k.GetValidatorSet(ctx)

	// return get proposer
	return validatorSet.GetProposer()
}

// SetValidatorIDToSignerAddr sets mapping for validator ID to signer address
func (k *Keeper) SetValidatorIDToSignerAddr(ctx sdk.Context, valID types.ValidatorID, signerAddr common.Address) {
	store := ctx.KVStore(k.StakingKey)
	store.Set(GetValidatorMapKey(valID.Bytes()), signerAddr.Bytes())
}

// GetSignerFromValidator get signer address from validator ID
func (k *Keeper) GetSignerFromValidatorID(ctx sdk.Context, valID types.ValidatorID) (common.Address, bool) {
	store := ctx.KVStore(k.StakingKey)
	key := GetValidatorMapKey(valID.Bytes())
	// check if validator address has been mapped
	if !store.Has(key) {
		return helper.ZeroAddress, false
	}
	// return address from bytes
	return common.BytesToAddress(store.Get(key)), true
}

// GetValidatorFromValAddr returns signer from validator ID
func (k *Keeper) GetValidatorFromValID(ctx sdk.Context, valID types.ValidatorID) (validator types.Validator, ok bool) {
	signerAddr, ok := k.GetSignerFromValidatorID(ctx, valID)
	if !ok {
		return validator, ok
	}
	// query for validator signer address
	validator, err := k.GetValidatorInfo(ctx, signerAddr.Bytes())
	if err != nil {
		return validator, false
	}
	return validator, true
}

// set last updated at for a validator
func (k *Keeper) SetLastUpdated(ctx sdk.Context, valID types.ValidatorID, blckNum *big.Int) sdk.Error {
	// get validator
	validator, ok := k.GetValidatorFromValID(ctx, valID)
	if !ok {
		return ErrInvalidMsg(k.Codespace, "unable to fetch validator", "ID", valID)
	}
	// make sure  new block num > old
	if blckNum.Cmp(validator.LastUpdated) != 1 {
		return ErrOldTx(k.Codespace)
	}
	validator.LastUpdated = blckNum
	err := k.AddValidator(ctx, validator)
	if err != nil {
		StakingLogger.Debug("Unable to update last updated", "Error", err, "Validator", validator.String())
		return ErrInvalidMsg(k.Codespace, "unable to add validator", "ID", valID, "Error", err)
	}
	return nil
}

// get last updated at for validator
func (k *Keeper) GetLastUpdated(ctx sdk.Context, valID types.ValidatorID) (updatedAT *big.Int, found bool) {
	// get validator
	validator, ok := k.GetValidatorFromValID(ctx, valID)
	if !ok {
		return nil, false
	}
	return validator.LastUpdated, true
}
