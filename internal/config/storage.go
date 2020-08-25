package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func (c *Config) FileCheck(enable bool) {
	if _, err := os.OpenFile(c.Files.SelfPath, os.O_RDONLY|os.O_CREATE, 0600); err != nil {
		log.Fatal("could not create vault-context config file")
	}
	if !enable {
		return
	}
	if _, err := os.Stat(c.Files.VaultPath); !os.IsNotExist(err) {
		log.Debug("deleting existing .vault file")
		if err := os.Remove(c.Files.VaultPath); err != nil {
			log.Error("could not delete existing .vault file. Please remove manually and then try again.")
		}
	}
	if file, err := os.OpenFile(c.Files.VaultPath, os.O_WRONLY|os.O_CREATE, 0600); err == nil {
		var ex string
		if self, err := os.Executable(); err == nil {
			ex = self
		} else {
			log.Error(err)
			return
		}
		if _, err := file.WriteString(fmt.Sprintf("token_helper = \"%s\"", ex)); err != nil {
			log.Error(err)
		} else {
			log.Debugf("created token-helper file at %s to use %s", c.Files.VaultPath, ex)
		}
	} else {
		log.Fatal("could not create vault config file")
	}
}
