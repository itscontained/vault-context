package commands

import (
	"encoding/json"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/jedib0t/go-pretty/table"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var wide = false

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured contexts",
	Long:  `Show a table of configured vault contexts`,
	PreRun: func(cmd *cobra.Command, args []string) {
		err = cfg.Keyring.InitKeyring()
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		list("table")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&wide, "wide", "w", false, "Look up all details about each token")
}

func list(format string) (ctxList []string) {
	switch format {
	case "table":
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.Style().Options.DrawBorder = false
		t.Style().Options.SeparateColumns = false
		t.Style().Options.SeparateHeader = false
		if wide {
			wideList(t)
		} else {
			headers := table.Row{"Alias", "URL", "Namespace", "Token"}
			t.AppendHeader(headers)
			for _, v := range cfg.VaultEnvs {
				var row table.Row
				if token, err := cfg.Keyring.Get(v.URL); err != nil {
					row = table.Row{v.Alias, v.URL, v.Namespace, "-"}
				} else {
					row = table.Row{v.Alias, v.URL, v.Namespace, token.Token}
				}
				t.AppendRow(row)
			}
		}
		t.Render()
		return nil
	case "search":
		for _, context := range cfg.VaultEnvs {
			if context.Alias != "" {
				ctxList = append(ctxList, context.Alias)
			} else {
				ctxList = append(ctxList, context.URL)
			}
		}
		return
	}
	return nil
}

func wideList(t table.Writer) {
	headers := table.Row{"Alias", "URL", "Namespace", "Token", "Username", "TTL", "Renewable", "Policies"}
	t.AppendHeader(headers)
	for _, v := range cfg.VaultEnvs {
		row := table.Row{v.Alias, v.URL, v.Namespace}
		token, err := cfg.Keyring.Get(v.URL)
		if err != nil {
			row = append(row, "-")
			t.AppendRow(row)
			continue
		}
		vaultConfig := vault.DefaultConfig()
		vaultConfig.Timeout = 5 * time.Second
		vaultConfig.MaxRetries = 1
		vaultConfig.Address = v.URL
		errRow := table.Row{"error", "error", "error", "error"}
		client, err := vault.NewClient(vaultConfig)
		if err != nil {
			row = append(row, errRow...)
			t.AppendRow(row)
			continue
		}
		client.SetToken(token.Token)
		if v.Namespace != "" {
			client.SetNamespace(v.Namespace)
		}
		s, err := client.Auth().Token().LookupSelf()
		if err != nil {
			row = append(row, errRow...)
			t.AppendRow(row)
			continue
		}
		ttl, err := s.Data["ttl"].(json.Number).Int64()
		if err != nil {
			row = append(row, errRow...)
			t.AppendRow(row)
			continue
		}
		row = append(row, token.Token, s.Data["display_name"], time.Duration(ttl)*time.Second,
			s.Data["renewable"], s.Data["policies"])
		t.AppendRow(row)
	}
}
