package rest

import (
	"fmt"
	"github.com/gxchain/TCPNetwork/x/tcp"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"

	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/gorilla/mux"
)

const (
	restName = "tcp"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/contracts/{%s}", storeName, restName), queryContractCodeHandle(cdc, cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("custom/%s", storeName), transferHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("custom/%s/{%s}", storeName, restName), deployContractHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("custom/%s/code/{%s}", storeName, restName), execContractHandler(cdc, cliCtx)).Methods("POST")
}


type transferReq struct {
	BaseReq rest.BaseReq	`json:"base_req"`
	From    string			`json:"from"`
	To 		string			`json:"to"`
	Amount  string			`json:"amount"`
}

// queryContractHandle
func queryContractCodeHandle(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars[restName]
		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/code/%s", storeName, address), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// transferHandler
func transferHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req transferReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		// parse from address
		addrFrom, err := sdk.AccAddressFromBech32(req.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// parse to address
		addrTo, err := sdk.AccAddressFromBech32(req.To)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// parse amount
		coins, err := sdk.ParseCoins(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := tcp.NewMsgTransfer(addrFrom, addrTo, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
	}
}



// deployContractHandler
func deployContractHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO

	}
}

// execContractHandler
func execContractHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO

	}
}
