package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	key        string
	secret     string
	passPhrase string
	version    string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ParseEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file %w", err)
	}

	if key, ok := os.LookupEnv("KEY"); ok {
		c.key = key
	} else {
		return errors.New("not found KEY in environment")
	}

	if key, ok := os.LookupEnv("SECRET"); ok {
		c.secret = key
	} else {
		return errors.New("not found SECRET in environment")
	}

	if key, ok := os.LookupEnv("PASS_PHRASE"); ok {
		c.passPhrase = key
	} else {
		return errors.New("not found PASS_PHRASE in environment")
	}

	if key, ok := os.LookupEnv("VERSION"); ok {
		c.version = key
	} else {
		return errors.New("not found VERSION in environment")
	}

	return nil
}

func (c *Config) Key() string {
	return c.key
}

func (c *Config) Secret() string {
	return c.secret
}

func (c *Config) PassPhrase() string {
	return c.passPhrase
}

func (c *Config) Version() string {
	return c.version
}
