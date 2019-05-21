package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hot3246624/TCPNetwork/x/tcp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

const (
	//flagFrom     = "from"
	flagAddress  = "address"
	flagCode     = "code"
	flagCodeHash = "codehash"
)

// GetCmdContractDeploy is the CLI command for deploying contract
func GetCmdContractDeploy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [from_addr] [contract_addr] [contract_code] [contract_hash]",
		Short: "deploy contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// from := viper.GetString(flagFrom)
			constractAddress := viper.GetString(flagAddress)
			code := viper.GetString(flagCode)
			codeHash := viper.GetString(flagCodeHash)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

			// get to contract address
			CIDAddr, err := sdk.AccAddressFromBech32(constractAddress)
			if err != nil {
				return err
			}
			// contract address must not exist
			err = cliCtx.EnsureAccountExistsFromAddr(CIDAddr)
			if err != nil {
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

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	//cmd.Flags().StringP(flagFrom, "f", "", "from address")
	cmd.Flags().StringP(flagAddress, "a", "", "contract address")
	cmd.Flags().StringP(flagCode, "c", "", "contract code")
	cmd.Flags().StringP(flagCodeHash, "s", "", "contract code hash")

	//cmd.MarkFlagRequired(flagFrom)
	cmd.MarkFlagRequired(flagAddress)
	cmd.MarkFlagRequired(flagCode)
	cmd.MarkFlagRequired(flagCodeHash)


	return cmd
}

// GetCmdContractExec is the CLI command for deploying contract
func GetCmdContractExec(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "exec [name] [value]",
		Short: "exec contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := tcp.NewMsgContractExec(fromAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}
