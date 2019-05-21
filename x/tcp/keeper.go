package tcp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the tcp Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func (k Keeper)GetContract(ctx sdk.Context, addr sdk.Address) ConAccount {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr.Bytes())) {
		return ConAccount{}
	}
	var conA ConAccount
	bz := store.Get([]byte(addr.Bytes()))
	k.cdc.MustUnmarshalBinaryBare(bz, &conA)
	return conA
}



func (k Keeper)GetResult(ctx sdk.Context, caller sdk.Address, contractAddr sdk.Address) []byte {
	conA := k.GetContract(ctx, contractAddr)
	return conA.Result[caller.String()]
}


func (k Keeper)DeployContract(ctx sdk.Context, contractAddr sdk.AccAddress, contactCode []byte, contactHash []byte) bool {
	// if there is a contract exist, cannot deploy contract.
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(contractAddr.Bytes())) {
		return false
	}
	conAccount := NewTCPWithDeploy(contractAddr, contactCode, contactHash)
	store.Set(contractAddr.Bytes(), k.cdc.MustMarshalBinaryBare(conAccount))
	return true
}

func (k Keeper)SetContractState(ctx sdk.Context, contractAddr sdk.AccAddress, addr sdk.AccAddress, result []byte) bool {
	conA := k.GetContract(ctx, contractAddr)
	conA.Result[addr.String()] = result
	store := ctx.KVStore(k.storeKey)
	store.Set(conA.Account.Address.Bytes(), k.cdc.MustMarshalBinaryBare(conA))
	return true
}