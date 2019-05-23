package tcp

import (
	"fmt"
	"github.com/gxchain/TCPNetwork/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "tcp" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgTransfer:
			return handleMsgTransfer(ctx, keeper, msg)
		case MsgContractDeploy:
			return handleContractDeploy(ctx, keeper, msg)
		case MsgContractExec:
			return handleMsgContractExec(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized tcp Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to transfer
func handleMsgTransfer(ctx sdk.Context, keeper Keeper, msg MsgTransfer) sdk.Result {
	// transfer coins
	if !msg.Value.IsValid() {
		return sdk.ErrInsufficientCoins("invalid coins").Result()
	}
	// check fee
	if msg.Fee.AmountOf(types.appCoin).Int64() < minTransferFee {
		return sdk.ErrInsufficientCoins("insufficient fee").Result()
	}

	_, err := keeper.coinKeeper.SendCoins(ctx, msg.From, msg.To, msg.Value)
	if err != nil {
		return sdk.ErrInsufficientCoins("does not have enough coins").Result()
	}

	return sdk.Result{}
}

// Handle a message to deploy contract
func handleContractDeploy(ctx sdk.Context, keeper Keeper, msg MsgContractDeploy) sdk.Result {
	// store code
	if msg.Code == nil || msg.CID == nil {
		return sdk.ErrUnknownRequest("there is invalid contract or not exist").Result()
	}
	// check fee
	if msg.Fee.AmountOf(types.appCoin).Int64() < minDeployFee {
		return sdk.ErrInsufficientCoins("insufficient fee").Result()
	}

	err := keeper.DeployContract(ctx, msg.CID, msg.Code, msg.CodeHash)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{}
}

// Handle a message to exec contract
func handleMsgContractExec(ctx sdk.Context, keeper Keeper, msg MsgContractExec) sdk.Result {
	if msg.Fee.AmountOf(types.appCoin).Int64() <= minExecFee {
		return sdk.ErrInsufficientCoins("insufficient fee").Result()
	}

	keeper.SetContractState(ctx, msg.CID, msg.From, msg.ResultHash[:])
	return sdk.Result{Data: msg.ResultHash}
}
