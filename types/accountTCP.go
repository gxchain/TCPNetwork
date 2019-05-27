package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"fmt"
	"encoding/json"
)

type ConAccount struct {
	Account  auth.BaseAccount  `json:"account"`
	Code     []byte            `json:"code"`
	CodeHash []byte            `json:"code_hash"`
	Result   []byte 		   `json:"result"`
}

func NewTCPWithDeploy(CID sdk.AccAddress, contractCode []byte, codeHash []byte) ConAccount {
	//hash and struct
	account := auth.NewBaseAccountWithAddress(CID)
	account.SetSequence(0)
	initDic := map[string]string{"test":"test"}
	initBytes, err := json.Marshal(initDic)
	if err != nil {
		return ConAccount{}
	}
	return ConAccount{
		Account:  account,
		Code:     contractCode,
		CodeHash: codeHash,
		Result:   initBytes,
	}

}

func (ca *ConAccount)Add(account sdk.AccAddress,result []byte) bool {
	key := account.String()
	//if ca.Contains(key) {
	//	return false
	//}
	fmt.Println("=========test Add method start========, result", ca.Result)
	temp, err := ca.convertToMap()
	if err != nil {
		return false
	}
	temp[key] = string(result)
	ca.Result, err = json.Marshal(temp)
	if err != nil {
		return false
	}
	fmt.Println("=========test Add method end========")

	return true
}

func (ca *ConAccount)Remove(result []byte) {
	return
}

func (ca ConAccount)Contains(account string) bool {
	//_, ok := ca.Result[account]
	return true
}

func (ca ConAccount)String(result []byte) string {
	return ""
}

func (ca ConAccount)convertToMap() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(ca.Result, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
