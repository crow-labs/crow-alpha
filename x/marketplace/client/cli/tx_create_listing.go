package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/crow-labs/crow/x/marketplace/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-listing [title] [description] [min-price]",
		Short: "Broadcast message createListing",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argMinPrice := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateListing(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argMinPrice,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
