package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// executeCmd represents the execute command
var triggerExecuteCmd = &cobra.Command{
	Use:     "execute <uuid>",
	Short:   "Execute the trigger with the specified UUID",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"run"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		if err := b.ExecuteTrigger(ctx, args[0]); err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}
	},
}

func init() {
	triggerCmd.AddCommand(triggerExecuteCmd)
}
