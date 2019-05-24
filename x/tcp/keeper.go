package tcp

import (
	"fmt"
	"github.com/gxchain/TCPNetwork/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper
	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc        *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the tcp Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func (k Keeper) GetContract(ctx sdk.Context, addr sdk.AccAddress) types.ConAccount {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr.Bytes())) {
		return types.ConAccount{}
	}
	var conA types.ConAccount
	bz := store.Get([]byte(addr.Bytes()))
	k.cdc.MustUnmarshalBinaryBare(bz, &conA)
	return conA
}

func (k Keeper) GetResult(ctx sdk.Context, caller sdk.AccAddress, contractAddr sdk.AccAddress) []byte {
	conA := k.GetContract(ctx, contractAddr)
	return conA.Result[caller.String()]
}

func (k Keeper) DeployContract(ctx sdk.Context, contractAddr sdk.AccAddress, contactCode []byte, contactHash []byte) sdk.Error {
	// if there is a contract exist, cannot deploy contract.
	store := ctx.KVStore(k.storeKey)
	if store.Has([]byte(contractAddr.Bytes())) {
		return sdk.ErrInternal("contract address already exists")
	}
	conAccount := types.NewTCPWithDeploy(contractAddr, contactCode, contactHash)
	store.Set(contractAddr.Bytes(), k.cdc.MustMarshalBinaryBare(conAccount))

	fmt.Println("==========deploy contract start===========")
	fmt.Println("conAccount info:", conAccount)
	account := k.GetContract(ctx, contractAddr)
	fmt.Println("deploy contract:", account)
	fmt.Println("==========deploy contract start===========")
	return nil
}

func (k Keeper) SetContractState(ctx sdk.Context, contractAddr sdk.AccAddress, fromAddr sdk.AccAddress, resultHash []byte) bool {
	conA := k.GetContract(ctx, contractAddr)
	fmt.Println("==========execute contract start===========")
	fmt.Println("contract info:", contractAddr, conA, conA.Result)
	fmt.Println("==========execute contract end===========")
	fmt.Println("fromAddr", string(fromAddr), "resultHash", resultHash, "Result", conA.Result)

	// maybe something is wrong here
	//conA.Result[string(fromAddr)] = resultHash

	fmt.Println("fromAddr", fromAddr, "resultHash", resultHash, "Result", conA.Result)

	store := ctx.KVStore(k.storeKey)
	store.Set(contractAddr.Bytes(), k.cdc.MustMarshalBinaryBare(conA))

	return true
}
