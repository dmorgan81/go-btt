package cmd

import (
	"context"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var triggerDeleteCmd = &cobra.Command{
	Use:     "delete <uuid>",
	Aliases: []string{"rm"},
	Short:   "Delete the trigger with the specified UUID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		if err := b.DeleteTrigger(ctx, args[0]); err != nil {
			exit(cmd, err)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerDeleteCmd)
}
