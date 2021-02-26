package cmd

import (
	"github.com/spf13/cobra"
)

// triggerCmd represents the trigger command
var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Commands related to triggers",
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
