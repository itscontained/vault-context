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

type Keychain struct {
	BackendType string                 `mapstructure:"backend"`
	Keychain    keychainBackendConfig  `mapstructure:"keychain"`
	KDEWallet   kdeWalletBackendConfig `mapstructure:"kdewallet"`
}

type keychainBackendConfig struct {
	Keychain       string `mapstructure:"keychain_name"`
	Synchronizable bool   `mapstructure:"icloud"`
}

type kdeWalletBackendConfig struct {
	Keychain string `mapstructure:"keychain_name"`
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
	URL       string `mapstructure:"url"`
	Namespace string `mapstructure:"namespace"`
	Alias     string `mapstructure:"alias"`
	Token     storage.Token
}

var Config = Cfg{
	Files: Files{},
	Keychain: Keychain{
		BackendType: "",
		Keychain:    keychainBackendConfig{},
		KDEWallet:   kdeWalletBackendConfig{},
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
	Write()
	return nil
}
