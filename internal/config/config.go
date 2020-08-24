package config

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/itscontained/vault-context/internal/storage"
)

type Cfg struct {
	Default   string `mapstructure:"default"`
	Files     Files
	Keychain  Keychain `mapstructure:"keychain"`
	Storage   *storage.Storage
	VaultEnvs []VaultEnv `mapstructure:"contexts"`
}

type Files struct {
	Self      string
	SelfDir   string
	SelfPath  string
	Vault     string
	VaultPath string
	Home      string
}
type VaultEnv struct {
	URL       string        `mapstructure:"url"`
	Namespace string        `mapstructure:"namespace"`
	Alias     string        `mapstructure:"alias"`
	Token     storage.Token `yaml:"-"`
}

var Config = Cfg{
	Files: Files{},
	Keychain: Keychain{
		BackendType: "automatic",
		Keychain:    keychainBackendConfig{},
		KDEWallet:   kdeWalletBackendConfig{},
		SecretService: secretServiceBackendConfig{
			Collection: "vault-context",
		},
		Pass: passBackendConfig{
			Prefix: "vault-context",
		},
		WinCred: winCredBackendConfig{
			Prefix: "vault-context",
		},
		File: fileBackendConfig{},
	},
	VaultEnvs: make([]VaultEnv, 0),
}

func (c *Cfg) Exists(v VaultEnv) error {
	for _, env := range c.VaultEnvs {
		if env.URL == v.URL {
			log.Debug("found matching vault address")
			if v.Alias == "" {
				return errors.New("no alias set. limited to one token for instance")
			}
			if env.Alias == v.Alias {
				return errors.New("a context with this alias and vault address already exists")
			}
			return nil
		}
	}
	return nil
}

func (c *Cfg) Add(url, namespace, alias string) error {
	v := VaultEnv{
		URL:       url,
		Namespace: namespace,
		Alias:     alias,
	}
	if err := c.Exists(v); err == nil {
		c.VaultEnvs = append(c.VaultEnvs, v)
	} else {
		return err
	}
	return nil
}
