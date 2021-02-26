package cmd

import (
	"github.com/spf13/cobra"
)

// triggerCmd represents the trigger command
var triggerCmd = &cobra.Command{
	Use: "trigger",
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
