// Package config contains config struct.
package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	DB        `yaml:"database"`
	Kucoin    `yaml:"kucoin"`
	Logger    `yaml:"logger"`
	MlService `yaml:"ml-service"`
}

func New() (*Config, error) {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		return nil, fmt.Errorf("environment CONFIG_PATH not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w: file %s dose not exists", err, configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("")
	}

	return &cfg, nil
}
