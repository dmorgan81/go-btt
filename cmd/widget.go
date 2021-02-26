package cmd

import (
	"github.com/spf13/cobra"
)

// widgetCmd represents the widget command
var widgetCmd = &cobra.Command{
	Use:   "widget",
	Short: "Commands related to widgets",
}

func init() {
	rootCmd.AddCommand(widgetCmd)
}
