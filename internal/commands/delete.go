package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <context>",
	Short: "Delete a saved context",
	Long:  `Delete a saved context from local storage`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		err = cfg.Keyring.InitKeyring()
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Delete(args[0])
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		err := cfg.Write()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
