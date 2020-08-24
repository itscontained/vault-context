package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/itscontained/vault-context/internal/storage"
)

func (c *Cfg) TokenHelper(cmd string) {
	url := os.Getenv("VAULT_ADDR")
	if url == "" {
		log.Error("not in a vault context")
		return
	}
	for _, context := range c.VaultEnvs {
		if context.URL == url {
			switch cmd {
			case "get":
				if t, err := c.Storage.Get(url); err == nil {
					fmt.Print(t.Token)
					return
				}
				return
			case "store":
				if stdin, err := ioutil.ReadAll(os.Stdin); err == nil {
					token := storage.Token{
						VaultAddr: url,
						Token:     strings.TrimSuffix(string(stdin), "\n"),
					}
					if err := Config.Storage.Store(token); err != nil {
						log.Error(err)
					}
				} else {
					log.Error("could not read input")
				}
			case "erase":
				if err := Config.Storage.Erase(url); err != nil {
					log.Error(err)
				}
			}
		}
	}
	Write()
}
