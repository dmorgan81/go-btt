package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dmorgan81/go-btt/btt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var variableCmd = &cobra.Command{
	Use:   "variable <name> <value>",
	Short: "Get or set a variable from BTT",
	Long: `The number of arguments determines whether to get or set a variable.
One argument means get the variable with that name.
Two arguments means set the variable with that name to the supplied value.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))

		if len(args) == 1 {
			s, err := b.GetVariable(ctx, args[0], viper.GetBool("persistent"), viper.GetBool("number"))
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			} else {
				fmt.Println(s)
			}
		} else {
			if err := b.SetVariable(ctx, args[0], args[1], viper.GetBool("persistent"), viper.GetBool("number")); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(variableCmd)

	variableCmd.Flags().Bool("persistent", false, "is variable persistent (default false)")
	variableCmd.Flags().Bool("number", false, "is variable a number (default false)")

	viper.BindPFlags(variableCmd.Flags())
}
