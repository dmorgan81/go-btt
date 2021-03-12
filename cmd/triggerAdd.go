package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var triggerAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"new", "create"},
	Args:    cobra.NoArgs,
	Short:   "Add a new trigger with JSON from stdin",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		uuid, err := b.AddTrigger(ctx, os.Stdin)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		} else {
			fmt.Println(uuid)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerAddCmd)
}
