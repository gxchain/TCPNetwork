package tcp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryContractCode = "code"
)

// Query Result Payload for a
type QueryResContractCode struct {
	Value string `json:"value"`

}

// implement fmt.Stringer
func (r QueryResContractCode) String() string {
	return r.Value
}

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryContractCode:
			return queryContractCode(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown tcp query endpoint")
		}
	}
}

// nolint: unparam
func queryContractCode(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	// get contract address
	contractAddr, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return []byte{}, sdk.ErrUnknownRequest("invalid contract address")
	}

	// get contract code
	value := string(keeper.GetContract(ctx, contractAddr).Code)
	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not get contract code")
	}


	bz, error := codec.MarshalJSONIndent(keeper.cdc, QueryResContractCode{value})
	if error != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
