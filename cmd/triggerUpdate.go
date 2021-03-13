package cmd

import (
	"context"
	"os"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var triggerUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update the trigger with the specified UUID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		if err := b.UpdateTrigger(ctx, args[0], os.Stdin); err != nil {
			exit(cmd, err)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerUpdateCmd)
}
