package config

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *Config) Write() error {
	viper.Set("keyring.backend", c.Keyring.Backend)
	viper.Set("keyring.config.filedir", c.Keyring.Config.FileDir)
	viper.Set("keyring.config.kwalletfolder", c.Keyring.Config.KWalletFolder)
	viper.Set("keyring.config.passdir", c.Keyring.Config.PassDir)
	viper.Set("keyring.config.passcmd", c.Keyring.Config.PassCmd)
	viper.Set("keyring.config.passprefix", c.Keyring.Config.PassPrefix)
	viper.Set("keyring.config.wincredprefix", c.Keyring.Config.WinCredPrefix)
	viper.Set("keyring.config.keychainsynchronizable", c.Keyring.Config.KeychainSynchronizable)
	viper.Set("keyring.config.keychaintrustapplication", c.Keyring.Config.KeychainTrustApplication)
	viper.Set("keyring.config.keychainaccessiblewhenunlocked", c.Keyring.Config.KeychainAccessibleWhenUnlocked)
	viper.Set("contexts", c.VaultEnvs)
	if err := viper.WriteConfig(); err != nil {
		return errors.New("could not write config file")
	} else {
		log.Debug("wrote to config file")
	}
	return nil
}
