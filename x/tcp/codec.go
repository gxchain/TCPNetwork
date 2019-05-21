package tcp

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransfer{}, "tcp/transfer", nil)
	cdc.RegisterConcrete(MsgContractDeploy{}, "tcp/deploy", nil)
	cdc.RegisterConcrete(MsgContractExec{}, "tcp/exec", nil)
}
