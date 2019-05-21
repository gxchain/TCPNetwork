package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/hot3246624/TCPNetwork/x/tcp"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"

	app "github.com/hot3246624/TCPNetwork"

	tcpclient "github.com/hot3246624/TCPNetwork/x/tcp/client"
	tcprest "github.com/hot3246624/TCPNetwork/x/tcp/client/rest"
)

const (
	storeAcc = "acc"
	storeTCP = "tcp"

	flagFrom     = "from"
	flagTo     = "to"
	flagAmount = "amount"
)


var defaultCLIHome = os.ExpandEnv("$HOME/.tcpcli")

func main() {

	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	mc := []sdk.ModuleClients{
		tcpclient.NewModuleClient(storeTCP, cdc),
	}

	rootCmd := &cobra.Command{
		Use:   "tcpcli",
		Short: "TCPNetwork Client",
	}

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(defaultCLIHome),
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
		transferCmd(cdc, mc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		keys.Commands(),
	)

	for _, m := range mc {
		rootCmd.AddCommand(m.GetTxCmd())
	}

	executor := cli.PrepareMainCmd(rootCmd, "TCP", defaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func registerRoutes(rs *lcd.RestServer) {
	rs.CliCtx = rs.CliCtx.WithAccountDecoder(rs.Cdc)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeAcc)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	tcprest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeTCP)
}

func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authcmd.GetAccountCmd(storeAcc, cdc),
	)

	for _, m := range mc {
		queryCmd.AddCommand(m.GetQueryCmd())
	}

	return queryCmd
}

func txCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		client.LineBreak,
	)

	return txCmd
}

func transferCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [from] [to] [amount]",
		Short: "transfer asset",
		Long: ` transfer asset from an address to another address.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))


			// from := viper.GetString(flagFrom)
			to := viper.GetString(flagTo)
			amount := viper.GetString(flagAmount)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

			// get to address
			toAddr, err := sdk.AccAddressFromBech32(to)
			if err != nil {
				return err
			}

			// get transfer amount
			coins, err := sdk.ParseCoins(amount)
			if err != nil {
				return err
			}

			// TODO
			msg := tcp.NewMsgTransfer(fromAddr, toAddr, coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().StringP(flagFrom, "f", "", "from address")
	cmd.Flags().StringP(flagTo, "t", "", "to address")
	cmd.Flags().StringP(flagAmount, "a", "", "coin amount")
	cmd.MarkFlagRequired(flagFrom)
	cmd.MarkFlagRequired(flagTo)
	cmd.MarkFlagRequired(flagAmount)



	return cmd
}


func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
