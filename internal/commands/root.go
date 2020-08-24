package commands

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/itscontained/vault-context/internal/config"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault-context",
	Short: "Vault context manager and token-helper",
	Long: `vault-context is a context manager and token helper designed to make
managing multiple vaults easier`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.Config.Init()
	viper.AddConfigPath(config.Config.Files.SelfDir)
	viper.SetConfigName(config.Config.Files.Self)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	config.Config.FileCheck(false)

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&config.Config); err != nil {
			log.Fatal("could not read config file")
		}
	} else {
		log.Fatal(err)
	}
}