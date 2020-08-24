package config

import log "github.com/sirupsen/logrus"

func (c *Cfg) Delete(ctx string) {
	for index, context := range c.VaultEnvs {
		if context.URL == ctx || context.Alias == ctx {
			c.VaultEnvs = append(c.VaultEnvs[:index], c.VaultEnvs[index+1:]...)
			if err := Config.Storage.Erase(context.URL); err != nil {
				log.Error(err)
			}
			log.Info("deleted context")
			return
		}
	}
	log.Error("could not find matching context")
}
