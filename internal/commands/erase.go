package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:    "erase",
	Hidden: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		err = cfg.Keyring.InitKeyring()
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg.TokenHelper("erase")
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
