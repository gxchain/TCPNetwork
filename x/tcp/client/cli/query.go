package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/gxchain/TCPNetwork/x/tcp"

)

// GetCmdCode queries a list of all names
func GetCmdCode(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "code",
		Short: "code",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			contractAddress := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/code/%s", queryRoute, contractAddress), nil)
			if err != nil {
				fmt.Printf("could not resolve contract - %s \n", string(contractAddress))
				return nil
			}

			var out tcp.QueryResContractCode
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
