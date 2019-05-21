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
	flagAddress  = "address"
	flagCode     = "code"
	flagCodeHash = "codehash"
)

// GetCmdContractDeploy is the CLI command for deploying contract
func GetCmdContractDeploy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy --from [from_addr] --address [contract_addr] --code [contract_code] --codehash [contract_hash]",
		Short: "deploy contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

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

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().StringP(flagAddress, "a", "", "contract address")
	cmd.Flags().StringP(flagCode, "c", "", "contract code")
	cmd.Flags().StringP(flagCodeHash, "s", "", "contract code hash")

	cmd.MarkFlagRequired(flagAddress)
	cmd.MarkFlagRequired(flagCode)
	cmd.MarkFlagRequired(flagCodeHash)

	return cmd
}

// GetCmdContractExec is the CLI command for deploying contract
func GetCmdContractExec(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec --from [from_address] --address [contract_address]",
		Short: "exec contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			constractAddress := viper.GetString(flagAddress)

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
				fmt.Println("contract address not exists")
				return err
			}

			msg := tcp.NewMsgContractExec(fromAddr, CIDAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().StringP(flagAddress, "a", "", "contract address")
	cmd.MarkFlagRequired(flagAddress)

	return cmd

}
