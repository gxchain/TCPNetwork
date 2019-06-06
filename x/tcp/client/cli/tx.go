package cli

import (
	"fmt"
	"strings"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gxchain/TCPNetwork/types"
	"github.com/gxchain/TCPNetwork/x/tcp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

const (
	flagContractAddress = "conAddress"
	flagCode            = "code"
	flagCodeHash        = "codeHash"
	flagTargets			= "targets"
	flagDataSources 	= "dataSources"
	flagCallerAddress   = "callAddress"
	flagState           = "state"
	flagProof           = "proof"
	flagResultHash      = "resultHash"
	flagPromise			= "promise"
)

// GetCmdContractDeploy is the CLI command for deploying contract
func GetCmdContractDeploy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy --conAddress [contract_addr] --code [contract_code] --codeHash [contract_hash] --targets [addr1,addr2] --dataSources [addr1,addr2]",
		Short: "deploy contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			contractAddress := viper.GetString(flagContractAddress)
			code := viper.GetString(flagCode)
			codeHash := viper.GetString(flagCodeHash)

			targetsStr := viper.GetString(flagTargets)
			dataSourcesStr := viper.GetString(flagDataSources)

			// parse target address
			targetAddressArr := []sdk.AccAddress{}
			targetAddressStr := strings.TrimSpace(targetsStr)
			addressStr := strings.Split(targetAddressStr, ",")
			for _, addrStr := range addressStr {
				addr, err := sdk.AccAddressFromBech32(addrStr)
				if err != nil {
					return err
				}
				targetAddressArr = append(targetAddressArr, addr)
			}

			// parse datasource address
			dataSourceAddressArr := []sdk.AccAddress{}
			dataSourceAddrStr := strings.TrimSpace(dataSourcesStr)
			addressStr = strings.Split(dataSourceAddrStr, ",")
			for _, addrStr := range addressStr {
				addr, err := sdk.AccAddressFromBech32(addrStr)
				if err != nil {
					return err
				}
				dataSourceAddressArr = append(dataSourceAddressArr, addr)
			}

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

			// get to contract address
			CIDAddr, err := sdk.AccAddressFromBech32(contractAddress)
			if err != nil {
				fmt.Println("invalid contract address")
				return err
			}
			// CIDAddress must not exist


			// TODO
			msg := tcp.NewMsgContractDeploy(fromAddr, CIDAddr, targetAddressArr, dataSourceAddressArr, []byte(code), []byte(codeHash))

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
	cmd.Flags().String(flagTargets, "", "target addresses")
	cmd.Flags().String(flagDataSources, "", "dataSource addresses")

	cmd.MarkFlagRequired(flagContractAddress)
	cmd.MarkFlagRequired(flagCode)
	cmd.MarkFlagRequired(flagCodeHash)
	cmd.MarkFlagRequired(flagTargets)
	cmd.MarkFlagRequired(flagDataSources)

	return cmd
}

// GetCmdContractExec is the CLI command for deploying contract
func GetCmdContractExec(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec --conAddress [contract_address] --callAddress [contract_caller] --state [state]  --proof [proof] --promise [promise] --resultHash [resultHash]",
		Short: "exec contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			contractAddress := viper.GetString(flagContractAddress)
			callerAddress := viper.GetString(flagCallerAddress)

			state := viper.GetString(flagState)
			proof := viper.GetString(flagProof)
			resultHash := viper.GetString(flagResultHash)
			promise := viper.GetString(flagPromise)

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
			callAddr, err := sdk.AccAddressFromBech32(callerAddress)
			if err != nil {
				return err
			}

			// TODO
			// contract address must not exist
			//err = cliCtx.EnsureAccountExistsFromAddr(CIDAddr)

			if err != nil {
				fmt.Println("contract address not exists")
				return err
			}

			req := types.RequestParam{
				From:        callAddr,
				CID:         CIDAddr,
				Proxy:       fromAddr,
				// DataSources: []types.Amount,
				Fee:         sdk.Coins{sdk.NewInt64Coin(types.AppCoin, 1)},
				Sig:         []byte(promise),
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
	cmd.Flags().String(flagPromise, "", "signature")

	cmd.MarkFlagRequired(flagContractAddress)
	cmd.MarkFlagRequired(flagCallerAddress)
	cmd.MarkFlagRequired(flagState)
	cmd.MarkFlagRequired(flagProof)
	cmd.MarkFlagRequired(flagResultHash)
	cmd.MarkFlagRequired(flagPromise)

	return cmd
}
