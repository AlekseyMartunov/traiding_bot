// Package kucoinorders allow you to place orders on the exchange.
package kucoinorders

import (
	"github.com/go-resty/resty/v2"
)

const (
	testEndpoint = "/api/v1/orders/test"
	endpoint     = "/api/v1/orders"
)

const (
	successfulCode = "200000"
)

type config interface {
	Key() string
	Secret() string
	Version() string
	PassPhrase() string
	BaseEndpoint() string
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Event(msg string, args ...any)
	Error(msg string, args ...any)
}

type KucoinOrderManager struct {
	log    logger
	cfg    config
	client *resty.Client
}

func NewKucoinOrderManager(l logger, c config) *KucoinOrderManager {
	return &KucoinOrderManager{
		log:    l,
		cfg:    c,
		client: resty.New(),
	}
}
