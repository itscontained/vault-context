package commands

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about your current context",
	Long:  ``,

}

func init() {
	rootCmd.AddCommand(infoCmd)
}