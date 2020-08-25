package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:    "store",
	Hidden: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		err = cfg.Keyring.InitKeyring()
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg.TokenHelper("store")
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
