package commands

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"
	"github.com/itscontained/vault-context/internal/utility"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:     "context",
	Short:   "Spawn a shell in the given context",
	Long:    `Spawn a shell in the given context. The shell is isolated from other shells.`,
	Aliases: []string{"ctx", "c"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			contexts := list("search")
			context := utility.FuzzyFind(contexts)
			shell(context)
		}
		shell(args[0])
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)

}

func shell(ctx string) {
	shell := os.Getenv("SHELL")
	env := syscall.Environ()
	var v config.VaultEnv
	for _, vaultEnv := range config.Config.VaultEnvs {
		if vaultEnv.Alias == ctx || vaultEnv.URL == ctx {
			v = vaultEnv
			break
		}
	}
	env = append(env, fmt.Sprintf("VAULT_ADDR=%s", v.URL), fmt.Sprintf("VAULT_CONTEXT_ADDR=%s", v.URL),
		"VAULT_CONTEXT=1")
	if v.Namespace != "" {
		env = append(env, fmt.Sprintf("VAULT_NAMESPACE=%s", v.Namespace))
	}
	_ = syscall.Exec(shell, []string{shell}, env)
}
