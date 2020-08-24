package commands

import (
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:    "store",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.TokenHelper("store")
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
