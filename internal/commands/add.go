package commands

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/utility"
)

var (
	url   = ""
	alias = ""
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <url> [alias]",
	Short: "Add a new Vault context",
	Long:  `Add a new vault context to your saved config with an optional alias`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a url argument")
		} else if len(args) > 2 {
			return errors.New("only a url and optional alias are allowed")
		}
		if !utility.IsUrl(args[0]) {
			return errors.New("invalid url")
		} else {
			url = args[0]
		}
		if len(args) == 2 {
			alias = args[1]
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		if err := cfg.Add(url, namespace, alias); err != nil {
			log.Error(err)
		} else {
			log.Info("context added")
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		err := cfg.Write()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("namespace", "n", "",
		"Namespace of the Vault Server (Enterprise). This will set the VAULT_NAMESPACE environment variable")
	addCmd.Flags().BoolP("default", "d", false,
		"Set this context as the default")
}
