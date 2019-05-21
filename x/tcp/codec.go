package tcp

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransfer{}, "tcp/Transfer", nil)
	cdc.RegisterConcrete(MsgContractDeploy{}, "tcp/ContractDeploy", nil)
	cdc.RegisterConcrete(MsgContractExec{}, "tcp/ContractExec", nil)
}
