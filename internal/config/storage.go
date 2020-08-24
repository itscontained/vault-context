package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/99designs/keyring"
	log "github.com/sirupsen/logrus"

	"github.com/itscontained/vault-context/internal/storage"
)

func (c *Cfg) FileCheck(enable bool) {
	if _, err := os.OpenFile(c.Files.SelfPath, os.O_RDONLY|os.O_CREATE, 0600); err != nil {
		log.Fatal("could not create vault-context config file")
	}
	if !enable {
		return
	}
	if file, err := os.OpenFile(c.Files.VaultPath, os.O_WRONLY|os.O_CREATE, 0600); err == nil {
		var ex string
		if self, err := os.Executable(); err == nil {
			ex = self
		} else {
			log.Error(err)
		}
		if _, err := file.WriteString(fmt.Sprintf("token_helper = \"%s\"", ex)); err != nil {
			log.Error(err)
		}
	} else {
		log.Fatal("could not create vault config file")
	}
}

func (c *Cfg) Init() {
	var err error
	f := &Files{
		Self:  "vault-context",
		Vault: ".vault",
	}
	if cUser, err := user.Current(); err != nil {
		log.Fatal(err)
	} else {
		f.Home = cUser.HomeDir
		f.SelfDir = filepath.Join(f.Home, ".config")
		f.SelfPath = filepath.Join(f.SelfDir, f.Self)
		f.VaultPath = filepath.Join(f.Home, f.Vault)
	}
	c.Files = *f

	storageCfg := keyring.Config{
		ServiceName: "vault-context",

		// keychain (macos)
		KeychainTrustApplication: true,
		KeychainSynchronizable:   Config.Keychain.Keychain.Synchronizable,
	}

	switch Config.Keychain.BackendType {
	case "automatic", "":
		storageCfg.AllowedBackends = storage.Backends
	case "keychain":
		storageCfg.AllowedBackends = []keyring.BackendType{keyring.KeychainBackend}
	case "kdewallet":
		storageCfg.AllowedBackends = []keyring.BackendType{keyring.KWalletBackend}
	default:
		log.Errorf("Unknown backend '%s'", Config.Keychain.BackendType)
	}

	if c.Storage, err = storage.New(storageCfg); err != nil {
		log.Fatalf("Unable to initialize backend '%s'", Config.Keychain.BackendType)
	}
}
