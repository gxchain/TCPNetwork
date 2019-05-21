package tcp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


// MsgTransfer defines a transfer message
type MsgTransfer struct {
	From sdk.AccAddress
	To sdk.AccAddress
	Value sdk.Coins
	// State []byte // TODO
	// Fee sdk.Coin

	// validation
	// From balance >= Value + Fee
	// From balance - Value - Fee = NewState of From
	// To balance + Value = NewState of To
}


// MsgContractDeploy defines a ContractDeploy message
type MsgContractDeploy struct {
	From sdk.AccAddress
	CID sdk.AccAddress
	Code []byte
	CodeHash []byte
	State []byte // TODO
	Fee sdk.Coins
}

// MsgContractExec defines a ontractExec message
type MsgContractExec struct {
	From sdk.AccAddress
	CID sdk.AccAddress
	State []byte // TODO
	RequestParams []RequestParam // TODO
	Proof []byte // TODO
	ResultHash []byte
	Fee sdk.Coins
}



// NewMsgTransfer is a constructor function for MsgTransfer
func NewMsgTransfer(from sdk.AccAddress, to sdk.AccAddress, value sdk.Coins) MsgTransfer {

	return MsgTransfer{
		From :  from,
		To : to,
		Value : value,
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
	contractAcc := NewTCPWithDeploy(CID, code, codeHash)
	return MsgContractDeploy{
		from,
		contractAcc.Account.Address,
		contractAcc.Code,
		contractAcc.CodeHash,
		[]byte{0},
		sdk.Coins{sdk.NewInt64Coin(appCoin,10)},
	}
}


// Route should return the name of the module
func (msg MsgContractDeploy) Route() string { return "tcp" }

// Type should return the action
func (msg MsgContractDeploy) Type() string { return "contract deploy" }

// ValidateBasic runs stateless checks on the message
func (msg MsgContractDeploy) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}

	if msg.CID.Empty() || msg.Code == nil || msg.CodeHash == nil{
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
func NewMsgContractExec(from sdk.AccAddress) MsgContractExec {
	req := RequestParam{
		//Some requestParam information
	}
	var reqS []RequestParam
	reqS = append(reqS, req)
	return MsgContractExec{
		from,
		req.CID,
		[]byte{1},
		reqS,
		[]byte("proof"),
		[]byte("result"),
		sdk.Coins{sdk.NewInt64Coin(appCoin,1)},

	}
}


// Route should return the name of the module
func (msg MsgContractExec) Route() string { return "tcp" }

// Type should return the action
func (msg MsgContractExec) Type() string { return "contract execute" }

// ValidateBasic runs stateless checks on the message
func (msg MsgContractExec) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
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
