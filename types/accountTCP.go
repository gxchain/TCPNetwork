package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
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
	temp, err := ca.convertToMap()
	if err != nil {
		return false
	}
	temp[key] = string(result)
	ca.Result, err = json.Marshal(temp)
	if err != nil {
		return false
	}
	return true
}

func (ca *ConAccount)Remove(result []byte) {
	return
}

func (ca ConAccount)Contains(account string) bool {
	dataMap, err := ca.convertToMap()
	if err != nil {
		return false
	}
	_, ok := dataMap[account]
	return ok
}

func (ca ConAccount)String(result []byte) string {
	return ""
}

func (ca ConAccount)ExecResult(caller sdk.AccAddress) string {
	data, err := ca.convertToMap()
	if err != nil {
		return ""
	}
	result, _ := data[caller.String()]
	return result
}

func (ca ConAccount)convertToMap() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(ca.Result, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
