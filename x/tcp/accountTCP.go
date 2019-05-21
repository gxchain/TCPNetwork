package tcp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type ConAccount struct {
	Account auth.BaseAccount	`json:"account"`
	Code []byte					`json:"code"`
	CodeHash []byte				`json:"code_hash"`
	Result map[string][]byte	`json:"result"`
}

func NewTCPWithDeploy(CID sdk.AccAddress, contractCode []byte, codeHash []byte) ConAccount{
	//addr := caller
	//nonce := uint64(8)
	//b := make([]byte, 8)
	//binary.BigEndian.PutUint64(b, nonce)
	//cAddr := append(addr, b...)
	//contractAddr := sdk.AccAddress(cAddr)

	//hash and struct
	account := auth.NewBaseAccountWithAddress(CID)
	account.SetSequence(0)
	ContractAcc := ConAccount{
		Account:account,
		Code:contractCode,
		CodeHash: codeHash,
	}

	return ContractAcc
}