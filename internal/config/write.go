package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Write() {
	viper.Set("keychain", Config.Keychain)
	viper.Set("contexts", Config.VaultEnvs)
	if err := viper.WriteConfig(); err != nil {
		log.Error("could not write config file")
	}
}
