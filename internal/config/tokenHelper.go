package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/itscontained/vault-context/internal/storage"
)

func (c *Config) TokenHelper(cmd string) {
	url := os.Getenv("VAULT_ADDR")
	if url == "" {
		log.Error("not in a vault context")
		return
	}
	for _, context := range c.VaultEnvs {
		if context.URL == url {
			switch cmd {
			case "get":
				if t, err := c.Keyring.Get(url); err == nil {
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
					if err := c.Keyring.Store(token); err != nil {
						log.Error(err)
					}
				} else {
					log.Error("could not read input")
				}
			case "erase":
				if err := c.Keyring.Erase(url); err != nil {
					log.Error(err)
				}
			}
		}
	}
}
