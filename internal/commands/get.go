package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:    "get",
	Hidden: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		err = cfg.Keyring.InitKeyring()
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg.TokenHelper("get")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
