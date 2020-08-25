package config

import (
	"errors"
	"os/user"
	"path/filepath"

	"github.com/99designs/keyring"
	log "github.com/sirupsen/logrus"

	"github.com/itscontained/vault-context/internal/storage"
)

type Config struct {
	Default   string          `mapstructure:"default"`
	Files     Files           `yaml:"-"`
	Keyring   storage.Keyring `mapstructure:"keyring"`
	VaultEnvs []VaultEnv      `mapstructure:"contexts" yaml:"contexts"`
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

func New() *Config {
	c := &Config{
		Files: Files{
			Self:  "vault-context",
			Vault: ".vault",
		},
		Keyring: storage.Keyring{
			Backend: "automatic",
			Config: keyring.Config{
				ServiceName:                    "vault-context",
				KeychainTrustApplication:       true,
				KeychainSynchronizable:         true,
				KeychainAccessibleWhenUnlocked: false,
				LibSecretCollectionName:        "vault-context",
				PassPrefix:                     "vault-context",
				WinCredPrefix:                  "vault-context",
			},
		},
		VaultEnvs: make([]VaultEnv, 0),
	}
	if cUser, err := user.Current(); err == nil {
		c.Files.Home = cUser.HomeDir
		c.Files.SelfDir = filepath.Join(c.Files.Home, ".config")
		c.Files.SelfPath = filepath.Join(c.Files.SelfDir, c.Files.Self)
		c.Files.VaultPath = filepath.Join(c.Files.Home, c.Files.Vault)

	} else {
		log.Fatal(err)
	}
	return c
}

func (c *Config) Exists(v VaultEnv) error {
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

func (c *Config) Add(url, namespace, alias string) error {
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

func (c *Config) Delete(ctx string) {
	for index, context := range c.VaultEnvs {
		if context.URL == ctx || context.Alias == ctx {
			c.VaultEnvs = append(c.VaultEnvs[:index], c.VaultEnvs[index+1:]...)
			if err := c.Keyring.Erase(context.URL); err != nil {
				log.Error(err)
			}
			log.Info("deleted context")
			return
		}
	}
	log.Fatal("could not find matching context")
}
