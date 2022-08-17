package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/crow-labs/crow/x/whitelist/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWhitelistUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-user [address]",
		Short: "Broadcast message whitelistUser",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWhitelistUser(
				clientCtx.GetFromAddress().String(),
				argAddress,
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
