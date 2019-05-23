package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	AppCoin = "nametoken"
)

const (
	// AddrLen defines a valid address length
	AddrLen = 20

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = "tcp"
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = "tcppub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = "tcpvaloper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = "tcpsvaloperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = "tcpvalcons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = "tcpvalconspub"
)

type Amount struct {
	Address []sdk.AccAddress `json:"address"`
	Value   []sdk.Coin       `json:"value"`
}

// user request for ContractExec
type RequestParam struct {

	From        sdk.AccAddress	`json:"from"`
	CID         sdk.AccAddress	`json:"cid"`
	Proxy       sdk.AccAddress	`json:"proxy"`
	DataSources []Amount		`json:"dataSource"`
	Fee         sdk.Coins		`json:"fee"`
	Sig         []byte			`json:"signature"`
}

type Balance struct {
	Address  []sdk.AccAddress	`json:"address"`
	Value 	[]sdk.Coins			`json:"value"`
}

type State struct {
	balances []Balance `json:"balances"`
}
