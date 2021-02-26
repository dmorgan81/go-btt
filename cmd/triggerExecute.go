package cmd

import (
	"context"
	"fmt"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// executeCmd represents the execute command
var triggerExecuteCmd = &cobra.Command{
	Use:     "execute <uuid>",
	Short:   "Executes the trigger with the specified UUID",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"run"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		err := b.ExecuteTrigger(ctx, args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerExecuteCmd)
}
