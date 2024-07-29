// Package config contains config structs.
package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// closeConfig is used to prevent accidental changes to internal fields in the main program
type closeConfig struct {
	CloseDB        closeDB        `yaml:"database"`
	CloseKucoin    closeKucoin    `yaml:"kucoin"`
	CloseLogger    closeLogger    `yaml:"logger"`
	CloseMLService closeMLService `yaml:"ml-service"`
}

func (cc *closeConfig) toExternalStruct() *Config {
	return &Config{
		DB:        cc.CloseDB.toExternalStruct(),
		Kucoin:    cc.CloseKucoin.toExternalStruct(),
		Logger:    cc.CloseLogger.toExternalStruct(),
		MlService: cc.CloseMLService.toExternalStruct(),
	}
}

type Config struct {
	DB        `yaml:"database"`
	Kucoin    `yaml:"kucoin-client"`
	Logger    `yaml:"logger"`
	MlService `yaml:"ml-service"`
}

func New() (*Config, error) {
	configPath, ok := os.LookupEnv("CONFIG_PATH_KUCOIN")
	if !ok {
		return nil, fmt.Errorf("environment CONFIG_PATH_KUCOIN is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w: file %s dose not exists", err, configPath)
	}

	var cfg closeConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("read config file err: %w", err)
	}

	return cfg.toExternalStruct(), nil
}
