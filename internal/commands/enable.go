package commands

import (
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Set vault-context as the token-helper",
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.FileCheck(true)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
