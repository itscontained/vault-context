package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.0"
	date    = "1970-01-01T00:00:00Z"
	commit  = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get the version and build date",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s, Build Date: %v, Commit: %v\n", version, date, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
