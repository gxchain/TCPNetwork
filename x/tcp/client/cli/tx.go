package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gxchain/TCPNetwork/x/tcp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

const (
	flagContractAddress  = "conAddress"
	flagCode     = "code"
	flagCodeHash = "codeHash"
	flagCallerAddress = "calAddress"
	flagState = "state"
	flagProof = "proof"
	flagResultHash = "resultHash"
)

// GetCmdContractDeploy is the CLI command for deploying contract
func GetCmdContractDeploy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy --conAddress [contract_addr] --code [contract_code] --codeHash [contract_hash]",
		Short: "deploy contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			contractAddress := viper.GetString(flagContractAddress)
			code := viper.GetString(flagCode)
			codeHash := viper.GetString(flagCodeHash)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

			// get to contract address
			CIDAddr, err := sdk.AccAddressFromBech32(contractAddress)
			if err != nil {
				return err
			}
			// contract address must not exist
			//err = cliCtx.EnsureAccountExistsFromAddr(CIDAddr)

			if err != nil {
				return err
			} else {
				fmt.Println("contract address must not exist")
			}

			// TODO
			msg := tcp.NewMsgContractDeploy(fromAddr, CIDAddr, []byte(code), []byte(codeHash))

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().StringP(flagContractAddress, "a", "", "contract address")
	cmd.Flags().StringP(flagCode, "c", "", "contract code")
	cmd.Flags().StringP(flagCodeHash, "s", "", "contract code hash")

	cmd.MarkFlagRequired(flagContractAddress)
	cmd.MarkFlagRequired(flagCode)
	cmd.MarkFlagRequired(flagCodeHash)

	return cmd
}

// GetCmdContractExec is the CLI command for deploying contract
func GetCmdContractExec(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec --conAddress [contract_address] --calAddress [contract_caller] --state [state]  --proof [proof] --resultHash [resultHash]",
		Short: "exec contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			contractAddress := viper.GetString(flagContractAddress)
			callerAddress := viper.GetString(flagCallerAddress)

			state := viper.GetString(flagState)
			proof := viper.GetString(flagProof)
			resultHash := viper.GetString(flagResultHash)


			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

			// get to contract address
			CIDAddr, err := sdk.AccAddressFromBech32(contractAddress)
			if err != nil {
				return err
			}

			// get to contract address
			calAddr, err := sdk.AccAddressFromBech32(callerAddress)
			if err != nil {
				return err
			}

			// contract address must not exist
			//err = cliCtx.EnsureAccountExistsFromAddr(CIDAddr)

			if err != nil {
				fmt.Println("contract address not exists")
				return err
			}

			req := tcp.RequestParam{
				From:calAddr,
				CID:CIDAddr,
				Proxy:fromAddr,

			}

			msg := tcp.NewMsgContractExec(fromAddr, []byte(state), []byte(proof), []byte(resultHash), req)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}


	cmd.Flags().StringP(flagContractAddress, "a", "", "contract address")
	cmd.Flags().StringP(flagCallerAddress, "c", "", "contract caller address")
	cmd.Flags().StringP(flagState, "s", "", "contract state")
	cmd.Flags().StringP(flagProof, "p", "", "contract proof")
	cmd.Flags().StringP(flagResultHash, "r", "", "contract exec result hash")

	cmd.MarkFlagRequired(flagContractAddress)
	cmd.MarkFlagRequired(flagCallerAddress)
	cmd.MarkFlagRequired(flagState)
	cmd.MarkFlagRequired(flagProof)
	cmd.MarkFlagRequired(flagResultHash)


	return cmd

}
