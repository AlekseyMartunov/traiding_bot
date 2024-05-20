package config

import (
	"errors"
	"os"
)

type Config struct {
	key        string
	secret     string
	passPhrase string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ParseEnvironment() error {
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
