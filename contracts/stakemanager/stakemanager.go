// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

import (
	"math/big"
	"strings"

	ethereum "github.com/maticnetwork/bor"
	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/core/types"
	"github.com/maticnetwork/bor/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakemanagerABI is the input ABI used to generate the binding from.
const StakemanagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentValidatorSet\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auctionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposerToSignerRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkPointBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stakeRewards\",\"type\":\"bool\"}],\"name\":\"restake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"enumStakeManager.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unJail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAmountToFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockInterval\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorState\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"stakerCount\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailCheckpoints\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supportsHistory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dynasty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"}],\"name\":\"stakeFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"getStakerDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegatorWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"}],\"name\":\"confirmAuctionBid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"checkpointTx\",\"type\":\"bytes\"}],\"name\":\"challangeStateRootUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentValidatorSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerNFT\",\"outputs\":[{\"internalType\":\"contractStaker\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"updateMinLockInPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"checkTwoByThreeMajority\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"updateCheckPointBlockInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentValidatorSetTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLockInPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"updateValidatorState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EPOCH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUpForFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumSlashedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"unstakeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"updateCheckpointReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewardsLiquidated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumSlashedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferRewards\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailCheckpoints\",\"type\":\"uint256\"}],\"name\":\"jail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNSTAKE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"updateTotalRewardsLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"}],\"name\":\"updateDynastyValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRootChain\",\"type\":\"address\"}],\"name\":\"changeRootChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakerNFT\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousRootChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRootChain\",\"type\":\"address\"}],\"name\":\"RootChainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"SignerChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"ReStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitEpoch\",\"type\":\"uint256\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDynasty\",\"type\":\"uint256\"}],\"name\":\"DynastyValueChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"}],\"name\":\"RewardUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"StakeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionAmount\",\"type\":\"uint256\"}],\"name\":\"StartAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newValidatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldValidatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ConfirmAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TopUpFee\",\"type\":\"event\"}]"

// Stakemanager is an auto generated Go binding around an Ethereum contract.
type Stakemanager struct {
	StakemanagerCaller     // Read-only binding to the contract
	StakemanagerTransactor // Write-only binding to the contract
	StakemanagerFilterer   // Log filterer for contract events
}

// StakemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakemanagerSession struct {
	Contract     *Stakemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakemanagerCallerSession struct {
	Contract *StakemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakemanagerTransactorSession struct {
	Contract     *StakemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakemanagerRaw struct {
	Contract *Stakemanager // Generic contract binding to access the raw methods on
}

// StakemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakemanagerCallerRaw struct {
	Contract *StakemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakemanagerTransactorRaw struct {
	Contract *StakemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakemanager creates a new instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanager(address common.Address, backend bind.ContractBackend) (*Stakemanager, error) {
	contract, err := bindStakemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// NewStakemanagerCaller creates a new read-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerCaller(address common.Address, caller bind.ContractCaller) (*StakemanagerCaller, error) {
	contract, err := bindStakemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerCaller{contract: contract}, nil
}

// NewStakemanagerTransactor creates a new write-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakemanagerTransactor, error) {
	contract, err := bindStakemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTransactor{contract: contract}, nil
}

// NewStakemanagerFilterer creates a new log filterer instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakemanagerFilterer, error) {
	contract, err := bindStakemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakemanagerFilterer{contract: contract}, nil
}

// bindStakemanager binds a generic wrapper to an already deployed contract.
func bindStakemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.StakemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transact(opts, method, params...)
}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) EPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "EPOCH_LENGTH")
	return *ret0, err
}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) EPOCHLENGTH() (*big.Int, error) {
	return _Stakemanager.Contract.EPOCHLENGTH(&_Stakemanager.CallOpts)
}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) EPOCHLENGTH() (*big.Int, error) {
	return _Stakemanager.Contract.EPOCHLENGTH(&_Stakemanager.CallOpts)
}

// MINDEPOSITSIZE is a free data retrieval call binding the contract method 0x26c0817e.
//
// Solidity: function MIN_DEPOSIT_SIZE() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) MINDEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "MIN_DEPOSIT_SIZE")
	return *ret0, err
}

// MINDEPOSITSIZE is a free data retrieval call binding the contract method 0x26c0817e.
//
// Solidity: function MIN_DEPOSIT_SIZE() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) MINDEPOSITSIZE() (*big.Int, error) {
	return _Stakemanager.Contract.MINDEPOSITSIZE(&_Stakemanager.CallOpts)
}

// MINDEPOSITSIZE is a free data retrieval call binding the contract method 0x26c0817e.
//
// Solidity: function MIN_DEPOSIT_SIZE() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) MINDEPOSITSIZE() (*big.Int, error) {
	return _Stakemanager.Contract.MINDEPOSITSIZE(&_Stakemanager.CallOpts)
}

// UNSTAKEDELAY is a free data retrieval call binding the contract method 0xe35e5d84.
//
// Solidity: function UNSTAKE_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) UNSTAKEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "UNSTAKE_DELAY")
	return *ret0, err
}

// UNSTAKEDELAY is a free data retrieval call binding the contract method 0xe35e5d84.
//
// Solidity: function UNSTAKE_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) UNSTAKEDELAY() (*big.Int, error) {
	return _Stakemanager.Contract.UNSTAKEDELAY(&_Stakemanager.CallOpts)
}

// UNSTAKEDELAY is a free data retrieval call binding the contract method 0xe35e5d84.
//
// Solidity: function UNSTAKE_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) UNSTAKEDELAY() (*big.Int, error) {
	return _Stakemanager.Contract.UNSTAKEDELAY(&_Stakemanager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "WITHDRAWAL_DELAY")
	return *ret0, err
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Stakemanager.Contract.WITHDRAWALDELAY(&_Stakemanager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Stakemanager.Contract.WITHDRAWALDELAY(&_Stakemanager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() constant returns(bytes32)
func (_Stakemanager *StakemanagerCaller) AccountStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "accountStateRoot")
	return *ret0, err
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() constant returns(bytes32)
func (_Stakemanager *StakemanagerSession) AccountStateRoot() ([32]byte, error) {
	return _Stakemanager.Contract.AccountStateRoot(&_Stakemanager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() constant returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) AccountStateRoot() ([32]byte, error) {
	return _Stakemanager.Contract.AccountStateRoot(&_Stakemanager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) AuctionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "auctionPeriod")
	return *ret0, err
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) AuctionPeriod() (*big.Int, error) {
	return _Stakemanager.Contract.AuctionPeriod(&_Stakemanager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) AuctionPeriod() (*big.Int, error) {
	return _Stakemanager.Contract.AuctionPeriod(&_Stakemanager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) CheckPointBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "checkPointBlockInterval")
	return *ret0, err
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _Stakemanager.Contract.CheckPointBlockInterval(&_Stakemanager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _Stakemanager.Contract.CheckPointBlockInterval(&_Stakemanager.CallOpts)
}

// CheckTwoByThreeMajority is a free data retrieval call binding the contract method 0xa062d09e.
//
// Solidity: function checkTwoByThreeMajority(bytes32 voteHash, bytes sigs) constant returns(uint256, uint256)
func (_Stakemanager *StakemanagerCaller) CheckTwoByThreeMajority(opts *bind.CallOpts, voteHash [32]byte, sigs []byte) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Stakemanager.contract.Call(opts, out, "checkTwoByThreeMajority", voteHash, sigs)
	return *ret0, *ret1, err
}

// CheckTwoByThreeMajority is a free data retrieval call binding the contract method 0xa062d09e.
//
// Solidity: function checkTwoByThreeMajority(bytes32 voteHash, bytes sigs) constant returns(uint256, uint256)
func (_Stakemanager *StakemanagerSession) CheckTwoByThreeMajority(voteHash [32]byte, sigs []byte) (*big.Int, *big.Int, error) {
	return _Stakemanager.Contract.CheckTwoByThreeMajority(&_Stakemanager.CallOpts, voteHash, sigs)
}

// CheckTwoByThreeMajority is a free data retrieval call binding the contract method 0xa062d09e.
//
// Solidity: function checkTwoByThreeMajority(bytes32 voteHash, bytes sigs) constant returns(uint256, uint256)
func (_Stakemanager *StakemanagerCallerSession) CheckTwoByThreeMajority(voteHash [32]byte, sigs []byte) (*big.Int, *big.Int, error) {
	return _Stakemanager.Contract.CheckTwoByThreeMajority(&_Stakemanager.CallOpts, voteHash, sigs)
}

// CheckpointReward is a free data retrieval call binding the contract method 0xda9c5cb5.
//
// Solidity: function checkpointReward() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) CheckpointReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "checkpointReward")
	return *ret0, err
}

// CheckpointReward is a free data retrieval call binding the contract method 0xda9c5cb5.
//
// Solidity: function checkpointReward() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) CheckpointReward() (*big.Int, error) {
	return _Stakemanager.Contract.CheckpointReward(&_Stakemanager.CallOpts)
}

// CheckpointReward is a free data retrieval call binding the contract method 0xda9c5cb5.
//
// Solidity: function checkpointReward() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) CheckpointReward() (*big.Int, error) {
	return _Stakemanager.Contract.CheckpointReward(&_Stakemanager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) CurrentEpoch() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentEpoch(&_Stakemanager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentEpoch(&_Stakemanager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) CurrentValidatorSetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "currentValidatorSetSize")
	return *ret0, err
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentValidatorSetSize(&_Stakemanager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentValidatorSetSize(&_Stakemanager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) CurrentValidatorSetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "currentValidatorSetTotalStake")
	return *ret0, err
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentValidatorSetTotalStake(&_Stakemanager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _Stakemanager.Contract.CurrentValidatorSetTotalStake(&_Stakemanager.CallOpts)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) Dynasty(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "dynasty")
	return *ret0, err
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) Dynasty() (*big.Int, error) {
	return _Stakemanager.Contract.Dynasty(&_Stakemanager.CallOpts)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) Dynasty() (*big.Int, error) {
	return _Stakemanager.Contract.Dynasty(&_Stakemanager.CallOpts)
}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() constant returns(uint256[])
func (_Stakemanager *StakemanagerCaller) GetCurrentValidatorSet(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "getCurrentValidatorSet")
	return *ret0, err
}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() constant returns(uint256[])
func (_Stakemanager *StakemanagerSession) GetCurrentValidatorSet() ([]*big.Int, error) {
	return _Stakemanager.Contract.GetCurrentValidatorSet(&_Stakemanager.CallOpts)
}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() constant returns(uint256[])
func (_Stakemanager *StakemanagerCallerSession) GetCurrentValidatorSet() ([]*big.Int, error) {
	return _Stakemanager.Contract.GetCurrentValidatorSet(&_Stakemanager.CallOpts)
}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) constant returns(uint256, uint256, uint256, address, uint256)
func (_Stakemanager *StakemanagerCaller) GetStakerDetails(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, *big.Int, *big.Int, common.Address, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Stakemanager.contract.Call(opts, out, "getStakerDetails", validatorId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) constant returns(uint256, uint256, uint256, address, uint256)
func (_Stakemanager *StakemanagerSession) GetStakerDetails(validatorId *big.Int) (*big.Int, *big.Int, *big.Int, common.Address, *big.Int, error) {
	return _Stakemanager.Contract.GetStakerDetails(&_Stakemanager.CallOpts, validatorId)
}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) constant returns(uint256, uint256, uint256, address, uint256)
func (_Stakemanager *StakemanagerCallerSession) GetStakerDetails(validatorId *big.Int) (*big.Int, *big.Int, *big.Int, common.Address, *big.Int, error) {
	return _Stakemanager.Contract.GetStakerDetails(&_Stakemanager.CallOpts, validatorId)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) GetValidatorId(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "getValidatorId", user)
	return *ret0, err
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetValidatorId(&_Stakemanager.CallOpts, user)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetValidatorId(&_Stakemanager.CallOpts, user)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Stakemanager *StakemanagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Stakemanager *StakemanagerSession) IsOwner() (bool, error) {
	return _Stakemanager.Contract.IsOwner(&_Stakemanager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Stakemanager *StakemanagerCallerSession) IsOwner() (bool, error) {
	return _Stakemanager.Contract.IsOwner(&_Stakemanager.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) constant returns(bool)
func (_Stakemanager *StakemanagerCaller) IsValidator(opts *bind.CallOpts, validatorId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "isValidator", validatorId)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) constant returns(bool)
func (_Stakemanager *StakemanagerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _Stakemanager.Contract.IsValidator(&_Stakemanager.CallOpts, validatorId)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) constant returns(bool)
func (_Stakemanager *StakemanagerCallerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _Stakemanager.Contract.IsValidator(&_Stakemanager.CallOpts, validatorId)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Stakemanager *StakemanagerCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Stakemanager *StakemanagerSession) Locked() (bool, error) {
	return _Stakemanager.Contract.Locked(&_Stakemanager.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Stakemanager *StakemanagerCallerSession) Locked() (bool, error) {
	return _Stakemanager.Contract.Locked(&_Stakemanager.CallOpts)
}

// MinLockInPeriod is a free data retrieval call binding the contract method 0xa548c547.
//
// Solidity: function minLockInPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) MinLockInPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "minLockInPeriod")
	return *ret0, err
}

// MinLockInPeriod is a free data retrieval call binding the contract method 0xa548c547.
//
// Solidity: function minLockInPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) MinLockInPeriod() (*big.Int, error) {
	return _Stakemanager.Contract.MinLockInPeriod(&_Stakemanager.CallOpts)
}

// MinLockInPeriod is a free data retrieval call binding the contract method 0xa548c547.
//
// Solidity: function minLockInPeriod() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) MinLockInPeriod() (*big.Int, error) {
	return _Stakemanager.Contract.MinLockInPeriod(&_Stakemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stakemanager *StakemanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stakemanager *StakemanagerSession) Owner() (common.Address, error) {
	return _Stakemanager.Contract.Owner(&_Stakemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stakemanager *StakemanagerCallerSession) Owner() (common.Address, error) {
	return _Stakemanager.Contract.Owner(&_Stakemanager.CallOpts)
}

// ProposerToSignerRewards is a free data retrieval call binding the contract method 0x1b3fa7ff.
//
// Solidity: function proposerToSignerRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) ProposerToSignerRewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "proposerToSignerRewards")
	return *ret0, err
}

// ProposerToSignerRewards is a free data retrieval call binding the contract method 0x1b3fa7ff.
//
// Solidity: function proposerToSignerRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) ProposerToSignerRewards() (*big.Int, error) {
	return _Stakemanager.Contract.ProposerToSignerRewards(&_Stakemanager.CallOpts)
}

// ProposerToSignerRewards is a free data retrieval call binding the contract method 0x1b3fa7ff.
//
// Solidity: function proposerToSignerRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ProposerToSignerRewards() (*big.Int, error) {
	return _Stakemanager.Contract.ProposerToSignerRewards(&_Stakemanager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Stakemanager *StakemanagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Stakemanager *StakemanagerSession) Registry() (common.Address, error) {
	return _Stakemanager.Contract.Registry(&_Stakemanager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Stakemanager *StakemanagerCallerSession) Registry() (common.Address, error) {
	return _Stakemanager.Contract.Registry(&_Stakemanager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Stakemanager *StakemanagerCaller) RootChain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "rootChain")
	return *ret0, err
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Stakemanager *StakemanagerSession) RootChain() (common.Address, error) {
	return _Stakemanager.Contract.RootChain(&_Stakemanager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Stakemanager *StakemanagerCallerSession) RootChain() (common.Address, error) {
	return _Stakemanager.Contract.RootChain(&_Stakemanager.CallOpts)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) SignerToValidator(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "signerToValidator", arg0)
	return *ret0, err
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) constant returns(uint256)
func (_Stakemanager *StakemanagerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.SignerToValidator(&_Stakemanager.CallOpts, arg0)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.SignerToValidator(&_Stakemanager.CallOpts, arg0)
}

// StakerNFT is a free data retrieval call binding the contract method 0x881a9d37.
//
// Solidity: function stakerNFT() constant returns(address)
func (_Stakemanager *StakemanagerCaller) StakerNFT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "stakerNFT")
	return *ret0, err
}

// StakerNFT is a free data retrieval call binding the contract method 0x881a9d37.
//
// Solidity: function stakerNFT() constant returns(address)
func (_Stakemanager *StakemanagerSession) StakerNFT() (common.Address, error) {
	return _Stakemanager.Contract.StakerNFT(&_Stakemanager.CallOpts)
}

// StakerNFT is a free data retrieval call binding the contract method 0x881a9d37.
//
// Solidity: function stakerNFT() constant returns(address)
func (_Stakemanager *StakemanagerCallerSession) StakerNFT() (common.Address, error) {
	return _Stakemanager.Contract.StakerNFT(&_Stakemanager.CallOpts)
}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() constant returns(bool)
func (_Stakemanager *StakemanagerCaller) SupportsHistory(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "supportsHistory")
	return *ret0, err
}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() constant returns(bool)
func (_Stakemanager *StakemanagerSession) SupportsHistory() (bool, error) {
	return _Stakemanager.Contract.SupportsHistory(&_Stakemanager.CallOpts)
}

// SupportsHistory is a free data retrieval call binding the contract method 0x7033e4a6.
//
// Solidity: function supportsHistory() constant returns(bool)
func (_Stakemanager *StakemanagerCallerSession) SupportsHistory() (bool, error) {
	return _Stakemanager.Contract.SupportsHistory(&_Stakemanager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Stakemanager *StakemanagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Stakemanager *StakemanagerSession) Token() (common.Address, error) {
	return _Stakemanager.Contract.Token(&_Stakemanager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Stakemanager *StakemanagerCallerSession) Token() (common.Address, error) {
	return _Stakemanager.Contract.Token(&_Stakemanager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "totalRewards")
	return *ret0, err
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalRewards() (*big.Int, error) {
	return _Stakemanager.Contract.TotalRewards(&_Stakemanager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalRewards() (*big.Int, error) {
	return _Stakemanager.Contract.TotalRewards(&_Stakemanager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalRewardsLiquidated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "totalRewardsLiquidated")
	return *ret0, err
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _Stakemanager.Contract.TotalRewardsLiquidated(&_Stakemanager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _Stakemanager.Contract.TotalRewardsLiquidated(&_Stakemanager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalStaked() (*big.Int, error) {
	return _Stakemanager.Contract.TotalStaked(&_Stakemanager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalStaked() (*big.Int, error) {
	return _Stakemanager.Contract.TotalStaked(&_Stakemanager.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalStakedFor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "totalStakedFor", user)
	return *ret0, err
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.TotalStakedFor(&_Stakemanager.CallOpts, user)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.TotalStakedFor(&_Stakemanager.CallOpts, user)
}

// ValAmountToFee is a free data retrieval call binding the contract method 0x40bb9cfb.
//
// Solidity: function valAmountToFee(uint256 ) constant returns(uint256 amount, uint256 fee)
func (_Stakemanager *StakemanagerCaller) ValAmountToFee(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	ret := new(struct {
		Amount *big.Int
		Fee    *big.Int
	})
	out := ret
	err := _Stakemanager.contract.Call(opts, out, "valAmountToFee", arg0)
	return *ret, err
}

// ValAmountToFee is a free data retrieval call binding the contract method 0x40bb9cfb.
//
// Solidity: function valAmountToFee(uint256 ) constant returns(uint256 amount, uint256 fee)
func (_Stakemanager *StakemanagerSession) ValAmountToFee(arg0 *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _Stakemanager.Contract.ValAmountToFee(&_Stakemanager.CallOpts, arg0)
}

// ValAmountToFee is a free data retrieval call binding the contract method 0x40bb9cfb.
//
// Solidity: function valAmountToFee(uint256 ) constant returns(uint256 amount, uint256 fee)
func (_Stakemanager *StakemanagerCallerSession) ValAmountToFee(arg0 *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _Stakemanager.Contract.ValAmountToFee(&_Stakemanager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) constant returns(uint256 amount, uint256 startEpoch, address user)
func (_Stakemanager *StakemanagerCaller) ValidatorAuction(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount     *big.Int
	StartEpoch *big.Int
	User       common.Address
}, error) {
	ret := new(struct {
		Amount     *big.Int
		StartEpoch *big.Int
		User       common.Address
	})
	out := ret
	err := _Stakemanager.contract.Call(opts, out, "validatorAuction", arg0)
	return *ret, err
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) constant returns(uint256 amount, uint256 startEpoch, address user)
func (_Stakemanager *StakemanagerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount     *big.Int
	StartEpoch *big.Int
	User       common.Address
}, error) {
	return _Stakemanager.Contract.ValidatorAuction(&_Stakemanager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) constant returns(uint256 amount, uint256 startEpoch, address user)
func (_Stakemanager *StakemanagerCallerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount     *big.Int
	StartEpoch *big.Int
	User       common.Address
}, error) {
	return _Stakemanager.Contract.ValidatorAuction(&_Stakemanager.CallOpts, arg0)
}

// ValidatorState is a free data retrieval call binding the contract method 0x5c248855.
//
// Solidity: function validatorState(uint256 ) constant returns(int256 amount, int256 stakerCount)
func (_Stakemanager *StakemanagerCaller) ValidatorState(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	ret := new(struct {
		Amount      *big.Int
		StakerCount *big.Int
	})
	out := ret
	err := _Stakemanager.contract.Call(opts, out, "validatorState", arg0)
	return *ret, err
}

// ValidatorState is a free data retrieval call binding the contract method 0x5c248855.
//
// Solidity: function validatorState(uint256 ) constant returns(int256 amount, int256 stakerCount)
func (_Stakemanager *StakemanagerSession) ValidatorState(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _Stakemanager.Contract.ValidatorState(&_Stakemanager.CallOpts, arg0)
}

// ValidatorState is a free data retrieval call binding the contract method 0x5c248855.
//
// Solidity: function validatorState(uint256 ) constant returns(int256 amount, int256 stakerCount)
func (_Stakemanager *StakemanagerCallerSession) ValidatorState(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _Stakemanager.Contract.ValidatorState(&_Stakemanager.CallOpts, arg0)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() constant returns(uint256)
func (_Stakemanager *StakemanagerCaller) ValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakemanager.contract.Call(opts, out, "validatorThreshold")
	return *ret0, err
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() constant returns(uint256)
func (_Stakemanager *StakemanagerSession) ValidatorThreshold() (*big.Int, error) {
	return _Stakemanager.Contract.ValidatorThreshold(&_Stakemanager.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() constant returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ValidatorThreshold() (*big.Int, error) {
	return _Stakemanager.Contract.ValidatorThreshold(&_Stakemanager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) constant returns(uint256 amount, uint256 reward, uint256 claimedRewards, uint256 slashedAmount, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, uint8 status)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ClaimedRewards    *big.Int
	SlashedAmount     *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	JailTime          *big.Int
	Signer            common.Address
	Status            uint8
}, error) {
	ret := new(struct {
		Amount            *big.Int
		Reward            *big.Int
		ClaimedRewards    *big.Int
		SlashedAmount     *big.Int
		ActivationEpoch   *big.Int
		DeactivationEpoch *big.Int
		JailTime          *big.Int
		Signer            common.Address
		Status            uint8
	})
	out := ret
	err := _Stakemanager.contract.Call(opts, out, "validators", arg0)
	return *ret, err
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) constant returns(uint256 amount, uint256 reward, uint256 claimedRewards, uint256 slashedAmount, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, uint8 status)
func (_Stakemanager *StakemanagerSession) Validators(arg0 *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ClaimedRewards    *big.Int
	SlashedAmount     *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	JailTime          *big.Int
	Signer            common.Address
	Status            uint8
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) constant returns(uint256 amount, uint256 reward, uint256 claimedRewards, uint256 slashedAmount, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, uint8 status)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ClaimedRewards    *big.Int
	SlashedAmount     *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	JailTime          *big.Int
	Signer            common.Address
	Status            uint8
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// ChallangeStateRootUpdate is a paid mutator transaction binding the contract method 0x7f5da718.
//
// Solidity: function challangeStateRootUpdate(bytes checkpointTx) returns()
func (_Stakemanager *StakemanagerTransactor) ChallangeStateRootUpdate(opts *bind.TransactOpts, checkpointTx []byte) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "challangeStateRootUpdate", checkpointTx)
}

// ChallangeStateRootUpdate is a paid mutator transaction binding the contract method 0x7f5da718.
//
// Solidity: function challangeStateRootUpdate(bytes checkpointTx) returns()
func (_Stakemanager *StakemanagerSession) ChallangeStateRootUpdate(checkpointTx []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.ChallangeStateRootUpdate(&_Stakemanager.TransactOpts, checkpointTx)
}

// ChallangeStateRootUpdate is a paid mutator transaction binding the contract method 0x7f5da718.
//
// Solidity: function challangeStateRootUpdate(bytes checkpointTx) returns()
func (_Stakemanager *StakemanagerTransactorSession) ChallangeStateRootUpdate(checkpointTx []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.ChallangeStateRootUpdate(&_Stakemanager.TransactOpts, checkpointTx)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_Stakemanager *StakemanagerTransactor) ChangeRootChain(opts *bind.TransactOpts, newRootChain common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "changeRootChain", newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_Stakemanager *StakemanagerSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ChangeRootChain(&_Stakemanager.TransactOpts, newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_Stakemanager *StakemanagerTransactorSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ChangeRootChain(&_Stakemanager.TransactOpts, newRootChain)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x5aac52f5.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, bytes sigs) returns(uint256)
func (_Stakemanager *StakemanagerTransactor) CheckSignatures(opts *bind.TransactOpts, blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, sigs []byte) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "checkSignatures", blockInterval, voteHash, stateRoot, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x5aac52f5.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, bytes sigs) returns(uint256)
func (_Stakemanager *StakemanagerSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, sigs []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.CheckSignatures(&_Stakemanager.TransactOpts, blockInterval, voteHash, stateRoot, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x5aac52f5.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, bytes sigs) returns(uint256)
func (_Stakemanager *StakemanagerTransactorSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, sigs []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.CheckSignatures(&_Stakemanager.TransactOpts, blockInterval, voteHash, stateRoot, sigs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xd02cfa56.
//
// Solidity: function claimRewards(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 index, bool transferRewards, bytes proof) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimRewards(opts *bind.TransactOpts, validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, index *big.Int, transferRewards bool, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimRewards", validatorId, accumBalance, accumSlashedAmount, index, transferRewards, proof)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xd02cfa56.
//
// Solidity: function claimRewards(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 index, bool transferRewards, bytes proof) returns()
func (_Stakemanager *StakemanagerSession) ClaimRewards(validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, index *big.Int, transferRewards bool, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, validatorId, accumBalance, accumSlashedAmount, index, transferRewards, proof)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xd02cfa56.
//
// Solidity: function claimRewards(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 index, bool transferRewards, bytes proof) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimRewards(validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, index *big.Int, transferRewards bool, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, validatorId, accumBalance, accumSlashedAmount, index, transferRewards, proof)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x7d4bd608.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactor) ConfirmAuctionBid(opts *bind.TransactOpts, validatorId *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "confirmAuctionBid", validatorId, commissionRate, signer, acceptDelegation)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x7d4bd608.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerSession) ConfirmAuctionBid(validatorId *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.ConfirmAuctionBid(&_Stakemanager.TransactOpts, validatorId, commissionRate, signer, acceptDelegation)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x7d4bd608.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactorSession) ConfirmAuctionBid(validatorId *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.ConfirmAuctionBid(&_Stakemanager.TransactOpts, validatorId, commissionRate, signer, acceptDelegation)
}

// DelegatorWithdrawal is a paid mutator transaction binding the contract method 0x7a0b19cd.
//
// Solidity: function delegatorWithdrawal(uint256 amount, address delegator) returns(bool)
func (_Stakemanager *StakemanagerTransactor) DelegatorWithdrawal(opts *bind.TransactOpts, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "delegatorWithdrawal", amount, delegator)
}

// DelegatorWithdrawal is a paid mutator transaction binding the contract method 0x7a0b19cd.
//
// Solidity: function delegatorWithdrawal(uint256 amount, address delegator) returns(bool)
func (_Stakemanager *StakemanagerSession) DelegatorWithdrawal(amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegatorWithdrawal(&_Stakemanager.TransactOpts, amount, delegator)
}

// DelegatorWithdrawal is a paid mutator transaction binding the contract method 0x7a0b19cd.
//
// Solidity: function delegatorWithdrawal(uint256 amount, address delegator) returns(bool)
func (_Stakemanager *StakemanagerTransactorSession) DelegatorWithdrawal(amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegatorWithdrawal(&_Stakemanager.TransactOpts, amount, delegator)
}

// Jail is a paid mutator transaction binding the contract method 0xd2203766.
//
// Solidity: function jail(uint256 validatorId, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerTransactor) Jail(opts *bind.TransactOpts, validatorId *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "jail", validatorId, jailCheckpoints)
}

// Jail is a paid mutator transaction binding the contract method 0xd2203766.
//
// Solidity: function jail(uint256 validatorId, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerSession) Jail(validatorId *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Jail(&_Stakemanager.TransactOpts, validatorId, jailCheckpoints)
}

// Jail is a paid mutator transaction binding the contract method 0xd2203766.
//
// Solidity: function jail(uint256 validatorId, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerTransactorSession) Jail(validatorId *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Jail(&_Stakemanager.TransactOpts, validatorId, jailCheckpoints)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Stakemanager *StakemanagerTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Stakemanager *StakemanagerSession) Lock() (*types.Transaction, error) {
	return _Stakemanager.Contract.Lock(&_Stakemanager.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Stakemanager *StakemanagerTransactorSession) Lock() (*types.Transaction, error) {
	return _Stakemanager.Contract.Lock(&_Stakemanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakemanager *StakemanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakemanager *StakemanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stakemanager.Contract.RenounceOwnership(&_Stakemanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakemanager *StakemanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stakemanager.Contract.RenounceOwnership(&_Stakemanager.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_Stakemanager *StakemanagerTransactor) Restake(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "restake", validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_Stakemanager *StakemanagerSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.Restake(&_Stakemanager.TransactOpts, validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_Stakemanager *StakemanagerTransactorSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.Restake(&_Stakemanager.TransactOpts, validatorId, amount, stakeRewards)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Stakemanager *StakemanagerTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Stakemanager *StakemanagerSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetToken(&_Stakemanager.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetToken(&_Stakemanager.TransactOpts, _token)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorId, uint256 slashingRate, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerTransactor) Slash(opts *bind.TransactOpts, validatorId *big.Int, slashingRate *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "slash", validatorId, slashingRate, jailCheckpoints)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorId, uint256 slashingRate, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerSession) Slash(validatorId *big.Int, slashingRate *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, validatorId, slashingRate, jailCheckpoints)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorId, uint256 slashingRate, uint256 jailCheckpoints) returns()
func (_Stakemanager *StakemanagerTransactorSession) Slash(validatorId *big.Int, slashingRate *big.Int, jailCheckpoints *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, validatorId, slashingRate, jailCheckpoints)
}

// Stake is a paid mutator transaction binding the contract method 0x06254a9c.
//
// Solidity: function stake(uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stake", amount, commissionRate, signer, acceptDelegation)
}

// Stake is a paid mutator transaction binding the contract method 0x06254a9c.
//
// Solidity: function stake(uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerSession) Stake(amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, amount, commissionRate, signer, acceptDelegation)
}

// Stake is a paid mutator transaction binding the contract method 0x06254a9c.
//
// Solidity: function stake(uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactorSession) Stake(amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, amount, commissionRate, signer, acceptDelegation)
}

// StakeFor is a paid mutator transaction binding the contract method 0x76324ea5.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactor) StakeFor(opts *bind.TransactOpts, user common.Address, amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stakeFor", user, amount, commissionRate, signer, acceptDelegation)
}

// StakeFor is a paid mutator transaction binding the contract method 0x76324ea5.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerSession) StakeFor(user common.Address, amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakeFor(&_Stakemanager.TransactOpts, user, amount, commissionRate, signer, acceptDelegation)
}

// StakeFor is a paid mutator transaction binding the contract method 0x76324ea5.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 commissionRate, address signer, bool acceptDelegation) returns()
func (_Stakemanager *StakemanagerTransactorSession) StakeFor(user common.Address, amount *big.Int, commissionRate *big.Int, signer common.Address, acceptDelegation bool) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakeFor(&_Stakemanager.TransactOpts, user, amount, commissionRate, signer, acceptDelegation)
}

// StartAuction is a paid mutator transaction binding the contract method 0x4fee13fc.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) StartAuction(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "startAuction", validatorId, amount)
}

// StartAuction is a paid mutator transaction binding the contract method 0x4fee13fc.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) StartAuction(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.StartAuction(&_Stakemanager.TransactOpts, validatorId, amount)
}

// StartAuction is a paid mutator transaction binding the contract method 0x4fee13fc.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) StartAuction(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.StartAuction(&_Stakemanager.TransactOpts, validatorId, amount)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0xac586292.
//
// Solidity: function topUpForFee(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) TopUpForFee(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "topUpForFee", validatorId, amount)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0xac586292.
//
// Solidity: function topUpForFee(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) TopUpForFee(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.TopUpForFee(&_Stakemanager.TransactOpts, validatorId, amount)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0xac586292.
//
// Solidity: function topUpForFee(uint256 validatorId, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) TopUpForFee(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.TopUpForFee(&_Stakemanager.TransactOpts, validatorId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakemanager *StakemanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakemanager *StakemanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.TransferOwnership(&_Stakemanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakemanager *StakemanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.TransferOwnership(&_Stakemanager.TransactOpts, newOwner)
}

// UnJail is a paid mutator transaction binding the contract method 0x3d02455b.
//
// Solidity: function unJail(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactor) UnJail(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unJail", validatorId)
}

// UnJail is a paid mutator transaction binding the contract method 0x3d02455b.
//
// Solidity: function unJail(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerSession) UnJail(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnJail(&_Stakemanager.TransactOpts, validatorId)
}

// UnJail is a paid mutator transaction binding the contract method 0x3d02455b.
//
// Solidity: function unJail(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactorSession) UnJail(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnJail(&_Stakemanager.TransactOpts, validatorId)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Stakemanager *StakemanagerTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Stakemanager *StakemanagerSession) Unlock() (*types.Transaction, error) {
	return _Stakemanager.Contract.Unlock(&_Stakemanager.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Stakemanager *StakemanagerTransactorSession) Unlock() (*types.Transaction, error) {
	return _Stakemanager.Contract.Unlock(&_Stakemanager.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactor) Unstake(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstake", validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactorSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xb97e3959.
//
// Solidity: function unstakeClaim(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 fee, uint256 index, bytes proof) returns()
func (_Stakemanager *StakemanagerTransactor) UnstakeClaim(opts *bind.TransactOpts, validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, fee *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstakeClaim", validatorId, accumBalance, accumSlashedAmount, fee, index, proof)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xb97e3959.
//
// Solidity: function unstakeClaim(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 fee, uint256 index, bytes proof) returns()
func (_Stakemanager *StakemanagerSession) UnstakeClaim(validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, fee *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnstakeClaim(&_Stakemanager.TransactOpts, validatorId, accumBalance, accumSlashedAmount, fee, index, proof)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xb97e3959.
//
// Solidity: function unstakeClaim(uint256 validatorId, uint256 accumBalance, uint256 accumSlashedAmount, uint256 fee, uint256 index, bytes proof) returns()
func (_Stakemanager *StakemanagerTransactorSession) UnstakeClaim(validatorId *big.Int, accumBalance *big.Int, accumSlashedAmount *big.Int, fee *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnstakeClaim(&_Stakemanager.TransactOpts, validatorId, accumBalance, accumSlashedAmount, fee, index, proof)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateCheckPointBlockInterval(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateCheckPointBlockInterval", _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_Stakemanager *StakemanagerSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCheckPointBlockInterval(&_Stakemanager.TransactOpts, _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCheckPointBlockInterval(&_Stakemanager.TransactOpts, _blocks)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateCheckpointReward(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateCheckpointReward", newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_Stakemanager *StakemanagerSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCheckpointReward(&_Stakemanager.TransactOpts, newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCheckpointReward(&_Stakemanager.TransactOpts, newReward)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateDynastyValue(opts *bind.TransactOpts, newDynasty *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateDynastyValue", newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_Stakemanager *StakemanagerSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateDynastyValue(&_Stakemanager.TransactOpts, newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateDynastyValue(&_Stakemanager.TransactOpts, newDynasty)
}

// UpdateMinLockInPeriod is a paid mutator transaction binding the contract method 0x98ee773b.
//
// Solidity: function updateMinLockInPeriod(uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateMinLockInPeriod(opts *bind.TransactOpts, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateMinLockInPeriod", epochs)
}

// UpdateMinLockInPeriod is a paid mutator transaction binding the contract method 0x98ee773b.
//
// Solidity: function updateMinLockInPeriod(uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) UpdateMinLockInPeriod(epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateMinLockInPeriod(&_Stakemanager.TransactOpts, epochs)
}

// UpdateMinLockInPeriod is a paid mutator transaction binding the contract method 0x98ee773b.
//
// Solidity: function updateMinLockInPeriod(uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateMinLockInPeriod(epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateMinLockInPeriod(&_Stakemanager.TransactOpts, epochs)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f283a86.
//
// Solidity: function updateSigner(uint256 validatorId, address _signer) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateSigner(opts *bind.TransactOpts, validatorId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateSigner", validatorId, _signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f283a86.
//
// Solidity: function updateSigner(uint256 validatorId, address _signer) returns()
func (_Stakemanager *StakemanagerSession) UpdateSigner(validatorId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateSigner(&_Stakemanager.TransactOpts, validatorId, _signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f283a86.
//
// Solidity: function updateSigner(uint256 validatorId, address _signer) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateSigner(validatorId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateSigner(&_Stakemanager.TransactOpts, validatorId, _signer)
}

// UpdateTotalRewardsLiquidated is a paid mutator transaction binding the contract method 0xe4e369a1.
//
// Solidity: function updateTotalRewardsLiquidated(uint256 _reward) returns(bool)
func (_Stakemanager *StakemanagerTransactor) UpdateTotalRewardsLiquidated(opts *bind.TransactOpts, _reward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateTotalRewardsLiquidated", _reward)
}

// UpdateTotalRewardsLiquidated is a paid mutator transaction binding the contract method 0xe4e369a1.
//
// Solidity: function updateTotalRewardsLiquidated(uint256 _reward) returns(bool)
func (_Stakemanager *StakemanagerSession) UpdateTotalRewardsLiquidated(_reward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateTotalRewardsLiquidated(&_Stakemanager.TransactOpts, _reward)
}

// UpdateTotalRewardsLiquidated is a paid mutator transaction binding the contract method 0xe4e369a1.
//
// Solidity: function updateTotalRewardsLiquidated(uint256 _reward) returns(bool)
func (_Stakemanager *StakemanagerTransactorSession) UpdateTotalRewardsLiquidated(_reward *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateTotalRewardsLiquidated(&_Stakemanager.TransactOpts, _reward)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0xa968882f.
//
// Solidity: function updateValidatorState(uint256 validatorId, uint256 epoch, int256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateValidatorState(opts *bind.TransactOpts, validatorId *big.Int, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateValidatorState", validatorId, epoch, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0xa968882f.
//
// Solidity: function updateValidatorState(uint256 validatorId, uint256 epoch, int256 amount) returns()
func (_Stakemanager *StakemanagerSession) UpdateValidatorState(validatorId *big.Int, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorState(&_Stakemanager.TransactOpts, validatorId, epoch, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0xa968882f.
//
// Solidity: function updateValidatorState(uint256 validatorId, uint256 epoch, int256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateValidatorState(validatorId *big.Int, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorState(&_Stakemanager.TransactOpts, validatorId, epoch, amount)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateValidatorThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateValidatorThreshold", newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_Stakemanager *StakemanagerSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorThreshold(&_Stakemanager.TransactOpts, newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorThreshold(&_Stakemanager.TransactOpts, newThreshold)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactor) WithdrawRewards(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "withdrawRewards", validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawRewards(&_Stakemanager.TransactOpts, validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_Stakemanager *StakemanagerTransactorSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawRewards(&_Stakemanager.TransactOpts, validatorId)
}

// StakemanagerClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Stakemanager contract.
type StakemanagerClaimRewardsIterator struct {
	Event *StakemanagerClaimRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerClaimRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimRewards represents a ClaimRewards event raised by the Stakemanager contract.
type StakemanagerClaimRewards struct {
	ValidatorId *big.Int
	Amount      *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimRewards(opts *bind.FilterOpts, validatorId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (*StakemanagerClaimRewardsIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimRewards", validatorIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimRewardsIterator{contract: _Stakemanager.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimRewards, validatorId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimRewards", validatorIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimRewards)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimRewards is a log parse operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimRewards(log types.Log) (*StakemanagerClaimRewards, error) {
	event := new(StakemanagerClaimRewards)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerConfirmAuctionIterator is returned from FilterConfirmAuction and is used to iterate over the raw logs and unpacked data for ConfirmAuction events raised by the Stakemanager contract.
type StakemanagerConfirmAuctionIterator struct {
	Event *StakemanagerConfirmAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerConfirmAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerConfirmAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerConfirmAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerConfirmAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerConfirmAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerConfirmAuction represents a ConfirmAuction event raised by the Stakemanager contract.
type StakemanagerConfirmAuction struct {
	NewValidatorId *big.Int
	OldValidatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConfirmAuction is a free log retrieval operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) FilterConfirmAuction(opts *bind.FilterOpts, newValidatorId []*big.Int, oldValidatorId []*big.Int, amount []*big.Int) (*StakemanagerConfirmAuctionIterator, error) {

	var newValidatorIdRule []interface{}
	for _, newValidatorIdItem := range newValidatorId {
		newValidatorIdRule = append(newValidatorIdRule, newValidatorIdItem)
	}
	var oldValidatorIdRule []interface{}
	for _, oldValidatorIdItem := range oldValidatorId {
		oldValidatorIdRule = append(oldValidatorIdRule, oldValidatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ConfirmAuction", newValidatorIdRule, oldValidatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerConfirmAuctionIterator{contract: _Stakemanager.contract, event: "ConfirmAuction", logs: logs, sub: sub}, nil
}

// WatchConfirmAuction is a free log subscription operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) WatchConfirmAuction(opts *bind.WatchOpts, sink chan<- *StakemanagerConfirmAuction, newValidatorId []*big.Int, oldValidatorId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var newValidatorIdRule []interface{}
	for _, newValidatorIdItem := range newValidatorId {
		newValidatorIdRule = append(newValidatorIdRule, newValidatorIdItem)
	}
	var oldValidatorIdRule []interface{}
	for _, oldValidatorIdItem := range oldValidatorId {
		oldValidatorIdRule = append(oldValidatorIdRule, oldValidatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ConfirmAuction", newValidatorIdRule, oldValidatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerConfirmAuction)
				if err := _Stakemanager.contract.UnpackLog(event, "ConfirmAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmAuction is a log parse operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) ParseConfirmAuction(log types.Log) (*StakemanagerConfirmAuction, error) {
	event := new(StakemanagerConfirmAuction)
	if err := _Stakemanager.contract.UnpackLog(event, "ConfirmAuction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerDynastyValueChangeIterator is returned from FilterDynastyValueChange and is used to iterate over the raw logs and unpacked data for DynastyValueChange events raised by the Stakemanager contract.
type StakemanagerDynastyValueChangeIterator struct {
	Event *StakemanagerDynastyValueChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerDynastyValueChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDynastyValueChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerDynastyValueChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerDynastyValueChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDynastyValueChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDynastyValueChange represents a DynastyValueChange event raised by the Stakemanager contract.
type StakemanagerDynastyValueChange struct {
	NewDynasty *big.Int
	OldDynasty *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDynastyValueChange is a free log retrieval operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Stakemanager *StakemanagerFilterer) FilterDynastyValueChange(opts *bind.FilterOpts) (*StakemanagerDynastyValueChangeIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DynastyValueChange")
	if err != nil {
		return nil, err
	}
	return &StakemanagerDynastyValueChangeIterator{contract: _Stakemanager.contract, event: "DynastyValueChange", logs: logs, sub: sub}, nil
}

// WatchDynastyValueChange is a free log subscription operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Stakemanager *StakemanagerFilterer) WatchDynastyValueChange(opts *bind.WatchOpts, sink chan<- *StakemanagerDynastyValueChange) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DynastyValueChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDynastyValueChange)
				if err := _Stakemanager.contract.UnpackLog(event, "DynastyValueChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDynastyValueChange is a log parse operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Stakemanager *StakemanagerFilterer) ParseDynastyValueChange(log types.Log) (*StakemanagerDynastyValueChange, error) {
	event := new(StakemanagerDynastyValueChange)
	if err := _Stakemanager.contract.UnpackLog(event, "DynastyValueChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the Stakemanager contract.
type StakemanagerJailedIterator struct {
	Event *StakemanagerJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerJailed represents a Jailed event raised by the Stakemanager contract.
type StakemanagerJailed struct {
	ValidatorId *big.Int
	ExitEpoch   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0xa1735a3843d9467dd849a217582720a8af66b9d034712e0b21b0f5ece80670cd.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch)
func (_Stakemanager *StakemanagerFilterer) FilterJailed(opts *bind.FilterOpts, validatorId []*big.Int, exitEpoch []*big.Int) (*StakemanagerJailedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var exitEpochRule []interface{}
	for _, exitEpochItem := range exitEpoch {
		exitEpochRule = append(exitEpochRule, exitEpochItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Jailed", validatorIdRule, exitEpochRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerJailedIterator{contract: _Stakemanager.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0xa1735a3843d9467dd849a217582720a8af66b9d034712e0b21b0f5ece80670cd.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch)
func (_Stakemanager *StakemanagerFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakemanagerJailed, validatorId []*big.Int, exitEpoch []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var exitEpochRule []interface{}
	for _, exitEpochItem := range exitEpoch {
		exitEpochRule = append(exitEpochRule, exitEpochItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Jailed", validatorIdRule, exitEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerJailed)
				if err := _Stakemanager.contract.UnpackLog(event, "Jailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJailed is a log parse operation binding the contract event 0xa1735a3843d9467dd849a217582720a8af66b9d034712e0b21b0f5ece80670cd.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch)
func (_Stakemanager *StakemanagerFilterer) ParseJailed(log types.Log) (*StakemanagerJailed, error) {
	event := new(StakemanagerJailed)
	if err := _Stakemanager.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stakemanager contract.
type StakemanagerOwnershipTransferredIterator struct {
	Event *StakemanagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Stakemanager contract.
type StakemanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakemanager *StakemanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakemanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerOwnershipTransferredIterator{contract: _Stakemanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakemanager *StakemanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakemanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerOwnershipTransferred)
				if err := _Stakemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakemanager *StakemanagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakemanagerOwnershipTransferred, error) {
	event := new(StakemanagerOwnershipTransferred)
	if err := _Stakemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerReStakedIterator is returned from FilterReStaked and is used to iterate over the raw logs and unpacked data for ReStaked events raised by the Stakemanager contract.
type StakemanagerReStakedIterator struct {
	Event *StakemanagerReStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerReStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerReStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerReStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerReStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerReStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerReStaked represents a ReStaked event raised by the Stakemanager contract.
type StakemanagerReStaked struct {
	ValidatorId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReStaked is a free log retrieval operation binding the contract event 0x9cc0e589f20d3310eb2ad571b23529003bd46048d0d1af29277dcf0aa3c398ce.
//
// Solidity: event ReStaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) FilterReStaked(opts *bind.FilterOpts, validatorId []*big.Int) (*StakemanagerReStakedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ReStaked", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerReStakedIterator{contract: _Stakemanager.contract, event: "ReStaked", logs: logs, sub: sub}, nil
}

// WatchReStaked is a free log subscription operation binding the contract event 0x9cc0e589f20d3310eb2ad571b23529003bd46048d0d1af29277dcf0aa3c398ce.
//
// Solidity: event ReStaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) WatchReStaked(opts *bind.WatchOpts, sink chan<- *StakemanagerReStaked, validatorId []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ReStaked", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerReStaked)
				if err := _Stakemanager.contract.UnpackLog(event, "ReStaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReStaked is a log parse operation binding the contract event 0x9cc0e589f20d3310eb2ad571b23529003bd46048d0d1af29277dcf0aa3c398ce.
//
// Solidity: event ReStaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) ParseReStaked(log types.Log) (*StakemanagerReStaked, error) {
	event := new(StakemanagerReStaked)
	if err := _Stakemanager.contract.UnpackLog(event, "ReStaked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerRewardUpdateIterator is returned from FilterRewardUpdate and is used to iterate over the raw logs and unpacked data for RewardUpdate events raised by the Stakemanager contract.
type StakemanagerRewardUpdateIterator struct {
	Event *StakemanagerRewardUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerRewardUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRewardUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerRewardUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerRewardUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRewardUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRewardUpdate represents a RewardUpdate event raised by the Stakemanager contract.
type StakemanagerRewardUpdate struct {
	NewReward *big.Int
	OldReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdate is a free log retrieval operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Stakemanager *StakemanagerFilterer) FilterRewardUpdate(opts *bind.FilterOpts) (*StakemanagerRewardUpdateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerRewardUpdateIterator{contract: _Stakemanager.contract, event: "RewardUpdate", logs: logs, sub: sub}, nil
}

// WatchRewardUpdate is a free log subscription operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Stakemanager *StakemanagerFilterer) WatchRewardUpdate(opts *bind.WatchOpts, sink chan<- *StakemanagerRewardUpdate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRewardUpdate)
				if err := _Stakemanager.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardUpdate is a log parse operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Stakemanager *StakemanagerFilterer) ParseRewardUpdate(log types.Log) (*StakemanagerRewardUpdate, error) {
	event := new(StakemanagerRewardUpdate)
	if err := _Stakemanager.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerRootChainChangedIterator is returned from FilterRootChainChanged and is used to iterate over the raw logs and unpacked data for RootChainChanged events raised by the Stakemanager contract.
type StakemanagerRootChainChangedIterator struct {
	Event *StakemanagerRootChainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerRootChainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRootChainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerRootChainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerRootChainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRootChainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRootChainChanged represents a RootChainChanged event raised by the Stakemanager contract.
type StakemanagerRootChainChanged struct {
	PreviousRootChain common.Address
	NewRootChain      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRootChainChanged is a free log retrieval operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_Stakemanager *StakemanagerFilterer) FilterRootChainChanged(opts *bind.FilterOpts, previousRootChain []common.Address, newRootChain []common.Address) (*StakemanagerRootChainChangedIterator, error) {

	var previousRootChainRule []interface{}
	for _, previousRootChainItem := range previousRootChain {
		previousRootChainRule = append(previousRootChainRule, previousRootChainItem)
	}
	var newRootChainRule []interface{}
	for _, newRootChainItem := range newRootChain {
		newRootChainRule = append(newRootChainRule, newRootChainItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RootChainChanged", previousRootChainRule, newRootChainRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRootChainChangedIterator{contract: _Stakemanager.contract, event: "RootChainChanged", logs: logs, sub: sub}, nil
}

// WatchRootChainChanged is a free log subscription operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_Stakemanager *StakemanagerFilterer) WatchRootChainChanged(opts *bind.WatchOpts, sink chan<- *StakemanagerRootChainChanged, previousRootChain []common.Address, newRootChain []common.Address) (event.Subscription, error) {

	var previousRootChainRule []interface{}
	for _, previousRootChainItem := range previousRootChain {
		previousRootChainRule = append(previousRootChainRule, previousRootChainItem)
	}
	var newRootChainRule []interface{}
	for _, newRootChainItem := range newRootChain {
		newRootChainRule = append(newRootChainRule, newRootChainItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RootChainChanged", previousRootChainRule, newRootChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRootChainChanged)
				if err := _Stakemanager.contract.UnpackLog(event, "RootChainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRootChainChanged is a log parse operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_Stakemanager *StakemanagerFilterer) ParseRootChainChanged(log types.Log) (*StakemanagerRootChainChanged, error) {
	event := new(StakemanagerRootChainChanged)
	if err := _Stakemanager.contract.UnpackLog(event, "RootChainChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerSignerChangeIterator is returned from FilterSignerChange and is used to iterate over the raw logs and unpacked data for SignerChange events raised by the Stakemanager contract.
type StakemanagerSignerChangeIterator struct {
	Event *StakemanagerSignerChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerSignerChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSignerChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerSignerChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerSignerChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSignerChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSignerChange represents a SignerChange event raised by the Stakemanager contract.
type StakemanagerSignerChange struct {
	ValidatorId *big.Int
	OldSigner   common.Address
	NewSigner   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignerChange is a free log retrieval operation binding the contract event 0x7dfd3bad1e3cac97d3b89ff06d78394523c4f08fdee4daa71a59160003240c89.
//
// Solidity: event SignerChange(uint256 indexed validatorId, address indexed oldSigner, address indexed newSigner)
func (_Stakemanager *StakemanagerFilterer) FilterSignerChange(opts *bind.FilterOpts, validatorId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (*StakemanagerSignerChangeIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SignerChange", validatorIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSignerChangeIterator{contract: _Stakemanager.contract, event: "SignerChange", logs: logs, sub: sub}, nil
}

// WatchSignerChange is a free log subscription operation binding the contract event 0x7dfd3bad1e3cac97d3b89ff06d78394523c4f08fdee4daa71a59160003240c89.
//
// Solidity: event SignerChange(uint256 indexed validatorId, address indexed oldSigner, address indexed newSigner)
func (_Stakemanager *StakemanagerFilterer) WatchSignerChange(opts *bind.WatchOpts, sink chan<- *StakemanagerSignerChange, validatorId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SignerChange", validatorIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSignerChange)
				if err := _Stakemanager.contract.UnpackLog(event, "SignerChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChange is a log parse operation binding the contract event 0x7dfd3bad1e3cac97d3b89ff06d78394523c4f08fdee4daa71a59160003240c89.
//
// Solidity: event SignerChange(uint256 indexed validatorId, address indexed oldSigner, address indexed newSigner)
func (_Stakemanager *StakemanagerFilterer) ParseSignerChange(log types.Log) (*StakemanagerSignerChange, error) {
	event := new(StakemanagerSignerChange)
	if err := _Stakemanager.contract.UnpackLog(event, "SignerChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerStakeUpdateIterator is returned from FilterStakeUpdate and is used to iterate over the raw logs and unpacked data for StakeUpdate events raised by the Stakemanager contract.
type StakemanagerStakeUpdateIterator struct {
	Event *StakemanagerStakeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerStakeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStakeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerStakeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerStakeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStakeUpdate represents a StakeUpdate event raised by the Stakemanager contract.
type StakemanagerStakeUpdate struct {
	ValidatorId *big.Int
	OldAmount   *big.Int
	NewAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeUpdate is a free log retrieval operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed oldAmount, uint256 indexed newAmount)
func (_Stakemanager *StakemanagerFilterer) FilterStakeUpdate(opts *bind.FilterOpts, validatorId []*big.Int, oldAmount []*big.Int, newAmount []*big.Int) (*StakemanagerStakeUpdateIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "StakeUpdate", validatorIdRule, oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakeUpdateIterator{contract: _Stakemanager.contract, event: "StakeUpdate", logs: logs, sub: sub}, nil
}

// WatchStakeUpdate is a free log subscription operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed oldAmount, uint256 indexed newAmount)
func (_Stakemanager *StakemanagerFilterer) WatchStakeUpdate(opts *bind.WatchOpts, sink chan<- *StakemanagerStakeUpdate, validatorId []*big.Int, oldAmount []*big.Int, newAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "StakeUpdate", validatorIdRule, oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStakeUpdate)
				if err := _Stakemanager.contract.UnpackLog(event, "StakeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeUpdate is a log parse operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed oldAmount, uint256 indexed newAmount)
func (_Stakemanager *StakemanagerFilterer) ParseStakeUpdate(log types.Log) (*StakemanagerStakeUpdate, error) {
	event := new(StakemanagerStakeUpdate)
	if err := _Stakemanager.contract.UnpackLog(event, "StakeUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stakemanager contract.
type StakemanagerStakedIterator struct {
	Event *StakemanagerStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStaked represents a Staked event raised by the Stakemanager contract.
type StakemanagerStaked struct {
	Signer          common.Address
	ValidatorId     *big.Int
	ActivationEpoch *big.Int
	Amount          *big.Int
	Total           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9cfd25589d1eb8ad71e342a86a8524e83522e3936c0803048c08f6d9ad974f40.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 indexed activationEpoch, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) FilterStaked(opts *bind.FilterOpts, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (*StakemanagerStakedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakedIterator{contract: _Stakemanager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9cfd25589d1eb8ad71e342a86a8524e83522e3936c0803048c08f6d9ad974f40.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 indexed activationEpoch, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakemanagerStaked, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9cfd25589d1eb8ad71e342a86a8524e83522e3936c0803048c08f6d9ad974f40.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 indexed activationEpoch, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) ParseStaked(log types.Log) (*StakemanagerStaked, error) {
	event := new(StakemanagerStaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerStartAuctionIterator is returned from FilterStartAuction and is used to iterate over the raw logs and unpacked data for StartAuction events raised by the Stakemanager contract.
type StakemanagerStartAuctionIterator struct {
	Event *StakemanagerStartAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerStartAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStartAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerStartAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerStartAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStartAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStartAuction represents a StartAuction event raised by the Stakemanager contract.
type StakemanagerStartAuction struct {
	ValidatorId   *big.Int
	Amount        *big.Int
	AuctionAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStartAuction is a free log retrieval operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Stakemanager *StakemanagerFilterer) FilterStartAuction(opts *bind.FilterOpts, validatorId []*big.Int, amount []*big.Int, auctionAmount []*big.Int) (*StakemanagerStartAuctionIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var auctionAmountRule []interface{}
	for _, auctionAmountItem := range auctionAmount {
		auctionAmountRule = append(auctionAmountRule, auctionAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "StartAuction", validatorIdRule, amountRule, auctionAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStartAuctionIterator{contract: _Stakemanager.contract, event: "StartAuction", logs: logs, sub: sub}, nil
}

// WatchStartAuction is a free log subscription operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Stakemanager *StakemanagerFilterer) WatchStartAuction(opts *bind.WatchOpts, sink chan<- *StakemanagerStartAuction, validatorId []*big.Int, amount []*big.Int, auctionAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var auctionAmountRule []interface{}
	for _, auctionAmountItem := range auctionAmount {
		auctionAmountRule = append(auctionAmountRule, auctionAmountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "StartAuction", validatorIdRule, amountRule, auctionAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStartAuction)
				if err := _Stakemanager.contract.UnpackLog(event, "StartAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartAuction is a log parse operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Stakemanager *StakemanagerFilterer) ParseStartAuction(log types.Log) (*StakemanagerStartAuction, error) {
	event := new(StakemanagerStartAuction)
	if err := _Stakemanager.contract.UnpackLog(event, "StartAuction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerThresholdChangeIterator is returned from FilterThresholdChange and is used to iterate over the raw logs and unpacked data for ThresholdChange events raised by the Stakemanager contract.
type StakemanagerThresholdChangeIterator struct {
	Event *StakemanagerThresholdChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerThresholdChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerThresholdChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerThresholdChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerThresholdChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerThresholdChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerThresholdChange represents a ThresholdChange event raised by the Stakemanager contract.
type StakemanagerThresholdChange struct {
	NewThreshold *big.Int
	OldThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChange is a free log retrieval operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Stakemanager *StakemanagerFilterer) FilterThresholdChange(opts *bind.FilterOpts) (*StakemanagerThresholdChangeIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ThresholdChange")
	if err != nil {
		return nil, err
	}
	return &StakemanagerThresholdChangeIterator{contract: _Stakemanager.contract, event: "ThresholdChange", logs: logs, sub: sub}, nil
}

// WatchThresholdChange is a free log subscription operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Stakemanager *StakemanagerFilterer) WatchThresholdChange(opts *bind.WatchOpts, sink chan<- *StakemanagerThresholdChange) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ThresholdChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerThresholdChange)
				if err := _Stakemanager.contract.UnpackLog(event, "ThresholdChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdChange is a log parse operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Stakemanager *StakemanagerFilterer) ParseThresholdChange(log types.Log) (*StakemanagerThresholdChange, error) {
	event := new(StakemanagerThresholdChange)
	if err := _Stakemanager.contract.UnpackLog(event, "ThresholdChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerTopUpFeeIterator is returned from FilterTopUpFee and is used to iterate over the raw logs and unpacked data for TopUpFee events raised by the Stakemanager contract.
type StakemanagerTopUpFeeIterator struct {
	Event *StakemanagerTopUpFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerTopUpFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerTopUpFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerTopUpFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerTopUpFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerTopUpFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerTopUpFee represents a TopUpFee event raised by the Stakemanager contract.
type StakemanagerTopUpFee struct {
	ValidatorId *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTopUpFee is a free log retrieval operation binding the contract event 0xaa00896a0cd7a8e4baccfe087fc5c9edc14be99df96ae50aa2e4b427c132e2c3.
//
// Solidity: event TopUpFee(uint256 indexed validatorId, uint256 indexed fee)
func (_Stakemanager *StakemanagerFilterer) FilterTopUpFee(opts *bind.FilterOpts, validatorId []*big.Int, fee []*big.Int) (*StakemanagerTopUpFeeIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "TopUpFee", validatorIdRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTopUpFeeIterator{contract: _Stakemanager.contract, event: "TopUpFee", logs: logs, sub: sub}, nil
}

// WatchTopUpFee is a free log subscription operation binding the contract event 0xaa00896a0cd7a8e4baccfe087fc5c9edc14be99df96ae50aa2e4b427c132e2c3.
//
// Solidity: event TopUpFee(uint256 indexed validatorId, uint256 indexed fee)
func (_Stakemanager *StakemanagerFilterer) WatchTopUpFee(opts *bind.WatchOpts, sink chan<- *StakemanagerTopUpFee, validatorId []*big.Int, fee []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "TopUpFee", validatorIdRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerTopUpFee)
				if err := _Stakemanager.contract.UnpackLog(event, "TopUpFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTopUpFee is a log parse operation binding the contract event 0xaa00896a0cd7a8e4baccfe087fc5c9edc14be99df96ae50aa2e4b427c132e2c3.
//
// Solidity: event TopUpFee(uint256 indexed validatorId, uint256 indexed fee)
func (_Stakemanager *StakemanagerFilterer) ParseTopUpFee(log types.Log) (*StakemanagerTopUpFee, error) {
	event := new(StakemanagerTopUpFee)
	if err := _Stakemanager.contract.UnpackLog(event, "TopUpFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerUnstakeInitIterator is returned from FilterUnstakeInit and is used to iterate over the raw logs and unpacked data for UnstakeInit events raised by the Stakemanager contract.
type StakemanagerUnstakeInitIterator struct {
	Event *StakemanagerUnstakeInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerUnstakeInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnstakeInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerUnstakeInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerUnstakeInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnstakeInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnstakeInit represents a UnstakeInit event raised by the Stakemanager contract.
type StakemanagerUnstakeInit struct {
	User              common.Address
	ValidatorId       *big.Int
	DeactivationEpoch *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeInit is a free log retrieval operation binding the contract event 0x51ff6d8ad000896362d1a3c9de9ad835adb6da92ec3ddee44018ef64f3c8b561.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 deactivationEpoch, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) FilterUnstakeInit(opts *bind.FilterOpts, user []common.Address, validatorId []*big.Int, amount []*big.Int) (*StakemanagerUnstakeInitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnstakeInitIterator{contract: _Stakemanager.contract, event: "UnstakeInit", logs: logs, sub: sub}, nil
}

// WatchUnstakeInit is a free log subscription operation binding the contract event 0x51ff6d8ad000896362d1a3c9de9ad835adb6da92ec3ddee44018ef64f3c8b561.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 deactivationEpoch, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) WatchUnstakeInit(opts *bind.WatchOpts, sink chan<- *StakemanagerUnstakeInit, user []common.Address, validatorId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnstakeInit)
				if err := _Stakemanager.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeInit is a log parse operation binding the contract event 0x51ff6d8ad000896362d1a3c9de9ad835adb6da92ec3ddee44018ef64f3c8b561.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 deactivationEpoch, uint256 indexed amount)
func (_Stakemanager *StakemanagerFilterer) ParseUnstakeInit(log types.Log) (*StakemanagerUnstakeInit, error) {
	event := new(StakemanagerUnstakeInit)
	if err := _Stakemanager.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakemanagerUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Stakemanager contract.
type StakemanagerUnstakedIterator struct {
	Event *StakemanagerUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnstaked represents a Unstaked event raised by the Stakemanager contract.
type StakemanagerUnstaked struct {
	User        common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address, validatorId []*big.Int) (*StakemanagerUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Unstaked", userRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnstakedIterator{contract: _Stakemanager.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakemanagerUnstaked, user []common.Address, validatorId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Unstaked", userRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnstaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Stakemanager *StakemanagerFilterer) ParseUnstaked(log types.Log) (*StakemanagerUnstaked, error) {
	event := new(StakemanagerUnstaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	return event, nil
}
