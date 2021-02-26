package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-btt",
	Short: "A utility for interacting with BetterTouchTool",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		lvl, err := log.ParseLevel(viper.GetString("log-level"))
		if err != nil {
			return err
		}
		log.SetLevel(lvl)

		if log.IsLevelEnabled(log.DebugLevel) {
			log.WithFields(viper.AllSettings()).Debug("settings")
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default %s/go-btt/config.yaml)", xdg.ConfigHome))
	rootCmd.PersistentFlags().String("log-level", "info", "")
	rootCmd.PersistentFlags().MarkHidden("log-level")
	rootCmd.PersistentFlags().String("addr", "http://localhost:50887", "BTT webserver address")
	rootCmd.PersistentFlags().String("secret", "", "shared secret for BTT")
	rootCmd.PersistentFlags().Duration("timeout", 10*time.Second, "timeout for requests")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(fmt.Sprintf("%s/go-btt", xdg.ConfigHome))
		viper.SetConfigName("config.yaml")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.WithField("config", viper.ConfigFileUsed()).Debug("using config")
	}
}
