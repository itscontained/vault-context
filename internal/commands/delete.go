package commands

import (
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <context>",
	Short: "Delete a saved context",
	Long:  `Delete a saved context from local storage`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.Config.Delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
