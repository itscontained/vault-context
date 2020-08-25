package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Get the current context URL",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		inContext := os.Getenv("VAULT_CONTEXT")
		vaultURL := os.Getenv("VAULT_CONTEXT_ADDR")
		if inContext != "" {
			fmt.Println(vaultURL)
		}
	},
}

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Get the current context namespace",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		inContext := os.Getenv("VAULT_CONTEXT")
		namespace := os.Getenv("VAULT_NAMESPACE")
		if inContext != "" {
			fmt.Println(namespace)
		}
	},
}

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Get the current context alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		inContext := os.Getenv("VAULT_CONTEXT")
		url := os.Getenv("VAULT_ADDR")
		if inContext != "" {
			for _, v := range cfg.VaultEnvs {
				if url == v.URL {
					fmt.Println(v.Alias)
					return
				}
			}
		}
	},
}

func init() {
	infoCmd.AddCommand(urlCmd)
	infoCmd.AddCommand(namespaceCmd)
	infoCmd.AddCommand(aliasCmd)
}
