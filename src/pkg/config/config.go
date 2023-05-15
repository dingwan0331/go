package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	Environment string `env:"ENVIRONMENT"`
	MysqlUrl    string `env:"MYSQL_URL"`
}

func NewConfig() Config {
	cfg := Config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		log.Fatalf("%+v", err)
	}

	return cfg
}
