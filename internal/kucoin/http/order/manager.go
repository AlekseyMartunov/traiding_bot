package kucoinorders

import (
	"github.com/go-resty/resty/v2"
)

const (
	testEndpoint       = "/api/v1/orders/test"
	endpoint           = "/api/v1/orders"
	symbolListEndpoint = "/api/v2/symbols"
	baseEndpoint       = "https://api.kucoin.com"
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

func (om *KucoinOrderManager) createHeaders(method, url, body string) map[string]string {
	return createSecretsHeaders(
		method,
		url,
		body,
		om.cfg.Secret(),
		om.cfg.PassPhrase(),
		om.cfg.Key(),
		om.cfg.Version(),
	)
}
