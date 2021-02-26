package cmd

import (
	"context"
	"fmt"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getTriggerCmd = &cobra.Command{
	Use:   "get <uuid>",
	Short: "Get a trigger by its UUID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		s, err := b.GetTrigger(ctx, args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(s)
		}
	},
}

func init() {
	triggerCmd.AddCommand(getTriggerCmd)
}
