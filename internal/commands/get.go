package commands

import (
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:    "get",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.TokenHelper("get")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
