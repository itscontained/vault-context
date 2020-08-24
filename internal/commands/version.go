package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version *string
	Date    *string
	Commit  *string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get the version and build date",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s, Build Date: %v, Commit: %v\n", *Version, *Date, *Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
