package cmd

import (
	"context"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var triggerNamedCmd = &cobra.Command{
	Use:   "named <name>",
	Short: "Trigger the specified named trigger",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		if err := b.Trigger(ctx, args[0], viper.GetBool("async")); err != nil {
			exit(cmd, err)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerNamedCmd)

	triggerNamedCmd.Flags().Bool("async", false, "exected named trigger asynchronously (default false)")

	viper.BindPFlags(triggerNamedCmd.Flags())
}
