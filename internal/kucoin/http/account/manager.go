// Package kucoinaccount allows you to make requests to the kucoin exchange
// to obtain info about currency accounts.
package kucoinaccount

import "github.com/go-resty/resty/v2"

const (
	baseEndpoint       = "https://api.kucoin.com"
	symbolListEndpoint = "/api/v2/symbols"
	accountEndpoint    = "/api/v1/accounts"
)

const (
	successfulCode = "200000"
)

type config interface {
	Key() string
	Secret() string
	Version() string
	PassPhrase() string
}

type logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
}

type AccountManager struct {
	log    logger
	cfg    config
	client *resty.Client
}

func New(l logger, c config) *AccountManager {
	return &AccountManager{
		log:    l,
		cfg:    c,
		client: resty.New(),
	}
}
