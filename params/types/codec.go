package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.NewLegacyAmino()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterCodec registers all necessary param module types with a given codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(ParameterChangeProposal{}, "heimdall/ParameterChangeProposal", nil)
}
