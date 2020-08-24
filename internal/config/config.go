package config

import (
	"errors"

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
	Token     storage.Token `mapstructure:"-"`
}

var Config = Cfg{
	Files: Files{},
	Keychain: Keychain{
		BackendType: "",
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

func (c *Cfg) Exists(url string) bool {
	for _, env := range c.VaultEnvs {
		if env.URL == url {
			return true
		}
	}
	return false
}

func (c *Cfg) Add(url, namespace, alias string) error {
	v := VaultEnv{
		URL:       url,
		Namespace: namespace,
		Alias:     alias,
	}
	if !c.Exists(url) {
		c.VaultEnvs = append(c.VaultEnvs, v)
	} else {
		return errors.New("context already exists")
	}
	return nil
}
