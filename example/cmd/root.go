package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dmorgan81/go-btt/btt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const bttVarName = "GoBTT-expl-uuid"

const triggerJson = `
{
	"BTTTouchBarButtonName" : "expl",
	"BTTTriggerType" : 629,
	"BTTTriggerTypeDescription" : "Touch Bar button",
	"BTTTriggerClass" : "BTTTriggerTypeTouchBar",
	"BTTPredefinedActionType" : -1,
	"BTTPredefinedActionName" : "No Action",
	"BTTEnabled2" : 1,
	"BTTRepeatDelay" : 0,
	"BTTNotesInsteadOfDescription" : 0,
	"BTTEnabled" : 1,
	"BTTModifierMode" : 0,
	"BTTOrder" : 6,
	"BTTDisplayOrder" : 0,
	"BTTMergeIntoTouchBarGroups" : 0,
	"BTTTriggerConfig" : {
	  "BTTTouchBarButtonColor" : "75.323769, 75.323769, 75.323769, 255.000000",
	  "BTTTouchBarItemIconWidth" : 22,
	  "BTTTouchBarButtonTextAlignment" : 0,
	  "BTTTouchBarButtonFontSize" : 15,
	  "BTTTouchBarAlternateBackgroundColor" : "75.323769, 75.323769, 75.323769, 255.000000",
	  "BTTTouchBarItemPlacement" : 0,
	  "BTTTouchBarAlwaysShowButton" : false,
	  "BTTTBWidgetWidth" : 400,
	  "BTTTouchBarIconTextOffset" : 5,
	  "BTTTouchBarButtonWidth" : 100,
	  "BTTTouchBarOnlyShowIcon" : false,
	  "BTTTouchBarFreeSpaceAfterButton" : 5,
	  "BTTTouchBarButtonName" : "expl",
	  "BTTTouchBarItemIconHeight" : 22,
	  "BTTTouchBarItemPadding" : 0
	}
  }
`

var rootCmd = &cobra.Command{
	Use: "btt-expl",
	Run: func(cmd *cobra.Command, args []string) {
		lvl, err := log.ParseLevel(viper.GetString("log-level"))
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}
		log.SetLevel(lvl)

		if log.IsLevelEnabled(log.DebugLevel) {
			log.WithFields(viper.AllSettings()).Debug("settings")
		}

		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		b := btt.New(viper.GetString("addr")).WithSecret(viper.GetString("secret"))

		uuid, err := b.GetStringVariable(ctx, bttVarName)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}

		if uuid == "" {
			buf := &bytes.Buffer{}
			if err := json.Compact(buf, []byte(triggerJson)); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}

			uuid, err = b.AddTrigger(ctx, bytes.NewReader(buf.Bytes()))
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}

			if err := b.SetPersistentStringVariable(ctx, bttVarName, uuid); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}
		}

		ticker := time.NewTicker(viper.GetDuration("interval"))
		defer ticker.Stop()

		done := make(chan os.Signal)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		var count uint32
		trigger := func() error {
			ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
			defer cancel()
			json := fmt.Sprintf(`{"BTTTouchBarButtonName":"expl %d"}`, count)
			return b.UpdateTrigger(ctx, uuid, strings.NewReader(json))
		}
		if err := trigger(); err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}

		for {
			select {
			case <-ticker.C:
				count += 1
				if err := trigger(); err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), err)
					os.Exit(1)
				}
			case <-done:
				return
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().String("log-level", "info", "")
	rootCmd.Flags().MarkHidden("log-level")
	rootCmd.Flags().String("addr", "http://localhost:50887", "BTT webserver address")
	rootCmd.Flags().String("secret", "", "shared secret for BTT")
	rootCmd.Flags().Duration("timeout", 10*time.Second, "timeout for requests")
	rootCmd.Flags().Duration("interval", 30*time.Second, "time between update operations")

	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
