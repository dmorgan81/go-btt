package cmd

import (
	"github.com/spf13/cobra"
)

var widgetCmd = &cobra.Command{
	Use:   "widget",
	Short: "Commands related to widgets",
}

func init() {
	rootCmd.AddCommand(widgetCmd)
}
