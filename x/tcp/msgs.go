package tcp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gxchain/TCPNetwork/types"
)

const (
	MinTransferFee       = 1
	MinContractDeployFee = 10
	MinContractExecFee   = 1
)


// message type and route constants
const (
	TypeMsgContractDeploy  = "tcp_deploy"
	RouteMsgContractDeploy = "tcp"

	TypeMsgContractExec  = "tcp_exec"
	RouteMsgContractExec = "tcp"
)

// MsgTransfer defines a transfer message
type MsgTransfer struct {
	From  sdk.AccAddress
	To    sdk.AccAddress
	Value sdk.Coins
	// State []byte // TODO
	Fee sdk.Coins
}

// MsgContractDeploy defines a ContractDeploy message
type MsgContractDeploy struct {
	From     sdk.AccAddress
	CID      sdk.AccAddress
	Code     []byte
	CodeHash []byte
	State    []byte // TODO
	Fee      sdk.Coins
}

// MsgContractExec defines a ontractExec message
type MsgContractExec struct {
	From         sdk.AccAddress
	CID          sdk.AccAddress
	State        []byte             // TODO
	RequestParam types.RequestParam // TODO
	Proof        []byte             // TODO
	ResultHash   []byte
	Fee          sdk.Coins
}

// NewMsgTransfer is a constructor function for MsgTransfer
func NewMsgTransfer(from sdk.AccAddress, to sdk.AccAddress, value sdk.Coins) MsgTransfer {
	return MsgTransfer{
		from,
		to,
		value,
		sdk.Coins{sdk.NewInt64Coin(types.AppCoin, MinTransferFee)},
	}
}

// Route should return the name of the module
func (msg MsgTransfer) Route() string { return "tcp" }

// Type should return the action
func (msg MsgTransfer) Type() string { return "transfer" }

// ValidateBasic runs stateless checks on the message
func (msg MsgTransfer) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}

	if msg.To.Empty() {
		return sdk.ErrInvalidAddress(msg.To.String())
	}

	if !msg.Value.IsAllPositive() {
		return sdk.ErrUnknownRequest("Transfer Value cannot be negative")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransfer) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// NewMsgContractDeploy is a constructor function for MsgTransfer
func NewMsgContractDeploy(from sdk.AccAddress, CID sdk.AccAddress, code []byte, codeHash []byte) MsgContractDeploy {
	// create contract account
	contractAcc := types.NewTCPWithDeploy(CID, code, codeHash)
	return MsgContractDeploy{
		from,
		contractAcc.Account.Address,
		contractAcc.Code,
		contractAcc.CodeHash,
		[]byte{0},
		sdk.Coins{sdk.NewInt64Coin(types.AppCoin, MinContractDeployFee)},
	}
}

// Route should return the name of the module
func (msg MsgContractDeploy) Route() string { return RouteMsgContractDeploy }

// Type should return the action
func (msg MsgContractDeploy) Type() string { return TypeMsgContractDeploy }

// ValidateBasic runs stateless checks on the message
func (msg MsgContractDeploy) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}

	if msg.CID.Empty() || msg.Code == nil || msg.CodeHash == nil {
		return sdk.ErrUnknownRequest("Contract cannot be nil")
	}

	if !msg.Fee.IsAllPositive() {
		return sdk.ErrUnknownRequest("Transfer Value cannot be negative")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgContractDeploy) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgContractDeploy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// NewMsgContractDeploy is a constructor function for MsgTransfer
func NewMsgContractExec(from sdk.AccAddress, state []byte, proof []byte, resultHash []byte, req types.RequestParam) MsgContractExec {
	return MsgContractExec{
		from,
		req.CID,
		state,
		req,
		proof,
		resultHash,
		sdk.Coins{sdk.NewInt64Coin(types.AppCoin, MinContractExecFee)}, //TODO set consume model
	}
}

// Route should return the name of the module
func (msg MsgContractExec) Route() string { return RouteMsgContractExec }

// Type should return the action
func (msg MsgContractExec) Type() string { return TypeMsgContractExec }

// ValidateBasic runs stateless checks on the message
func (msg MsgContractExec) ValidateBasic() sdk.Error {
	// check address
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}
	// check address
	if msg.RequestParam.From.Empty() {
		return sdk.ErrInvalidAddress(msg.RequestParam.From.String())
	}
	// check address
	if msg.RequestParam.Proxy.Empty() {
		return sdk.ErrInvalidAddress(msg.RequestParam.Proxy.String())
	}

	//TODO validate the CID whether is the right Contract.
	if msg.CID.Empty() {
		return sdk.ErrUnknownRequest("Contract cannot be nil")
	}

	if !msg.Fee.IsAllPositive() {
		return sdk.ErrUnknownRequest("Transfer Value cannot be negative")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgContractExec) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgContractExec) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}