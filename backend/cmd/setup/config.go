package setup

import (
	"github.com/esceer/vault/backend/cmd/config"
	"github.com/kelseyhightower/envconfig"
)

func Config() (*config.Config, error) {
	cfg := config.NewConfig()
	err := envconfig.Process("", cfg)
	return cfg, err
}
