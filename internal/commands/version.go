package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version   = "0.0.0"
	buildDate = "1970-01-01T00:00:00Z"
	commit    = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "get",
	Short: "get the version and build date",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s, Build Date: %s, Commit: %s", version, buildDate, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}