// Package kucoinaccount allows you to make requests to the kucoin exchange
// to obtain info about currency accounts.
package kucoinaccount

import "github.com/go-resty/resty/v2"

const (
	symbolListEndpoint = "/api/v2/symbols"
	accountEndpoint    = "/api/v1/accounts"
)

const (
	successfulCode = "200000"
)

type config interface {
	GetKey() string
	GetSecret() string
	GetVersion() string
	GetPassPhrase() string
	GetBaseEndpoint() string
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Event(msg string, args ...any)
	Error(msg string, args ...any)
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
