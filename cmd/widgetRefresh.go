package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var widgetRefreshCmd = &cobra.Command{
	Use:   "refresh <uuid>",
	Short: "Refresh the widget with the specified UUID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))
		if err := b.RefreshWidget(ctx, args[0]); err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}
	},
}

func init() {
	widgetCmd.AddCommand(widgetRefreshCmd)
}
