package tcp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


const (
	appCoin = "nametoken"
)

type Amount struct {
	Address []sdk.AccAddress	`json:"address"`
	Value   []sdk.Coin			`json:"value"`
}

// user request for ContractExec
type RequestParam struct {
	From        sdk.AccAddress	`json:"from"`
	CID         sdk.AccAddress	`json:"cid"`
	Proxy       sdk.AccAddress	`json:"proxy"`
	DataSources []Amount		`json:"datasource"`
	Fee         sdk.Coin		`json:"fee"`
	Sig         []byte			`json:"signature"`

}

type Balance struct {
	Address  []sdk.AccAddress	`json:"address"`
	Value []sdk.Coin			`json:"value"`
}

type State struct {
	balances []Balance			`json:"balances"`
}