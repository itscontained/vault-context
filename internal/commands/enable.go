package commands

import (
	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Set vault-context as the token-helper",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.FileCheck(true)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
