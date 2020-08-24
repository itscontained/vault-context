package commands

import (
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:    "erase",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.TokenHelper("erase")
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
