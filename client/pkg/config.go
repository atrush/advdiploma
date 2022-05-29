package pkg

import (
	"errors"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

//  Config stores server config params.
type Config struct {
	MasterKey   string `env:"MASTER_KEY"`
	SyncTimeout time.Duration
}

//  Default config params.
const (
	defMasterKey = "mytestmasterkey"
)

//  NewConfig inits new config.
//  Reads flag params over default params, then redefines  with environment params.
func NewConfig() (*Config, error) {
	cfg := Config{}

	cfg.readFlagConfig()
	if err := cfg.readEnvConfig(); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("ошибка инициализации конфига: %w", err)
	}
	return &cfg, nil
}

//  Validate validates config params.
func (c *Config) Validate() error {
	if len(c.MasterKey) == 0 {
		return errors.New("master key is empty")
	}
	return nil
}

//  readFlagConfig reads flag params over default params.
func (c *Config) readFlagConfig() {
	flagConfig := &Config{}
	flag.StringVar(&flagConfig.MasterKey, "m", defMasterKey, "master key")
	flag.Parse()

	c.redefineConfig(flagConfig)
}

//  redefineConfig redefines config with new config.
//  if string not empty override
//  if bool true override
func (c *Config) redefineConfig(nc *Config) {
	if nc.MasterKey != "" {
		c.MasterKey = nc.MasterKey
	}
}

//  readEnvConfig redefines config params with environment params.
func (c *Config) readEnvConfig() error {
	envConfig := &Config{}

	if err := env.Parse(envConfig); err != nil {
		return fmt.Errorf("ошибка чтения переменных окружения:%w", err)
	}

	c.redefineConfig(envConfig)
	return nil
}
