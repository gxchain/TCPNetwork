package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetCmdCode queries a list of all names
func GetCmdCode(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "names",
		Short: "names",
		// Args:  cobra.ExactArgs(1),
		// TODO
	}
}
