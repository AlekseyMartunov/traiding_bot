package kucoinaccount

import "github.com/go-resty/resty/v2"

const (
	baseEndpoint    = "https://api.kucoin.com"
	accountEndpoint = "/api/v1/accounts"
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

func NewAccountManager(l logger, c config) *AccountManager {
	return &AccountManager{
		log:    l,
		cfg:    c,
		client: resty.New(),
	}
}

func (am *AccountManager) createHeaders(method, url, body string) map[string]string {
	return createSecretsHeaders(
		method,
		url,
		body,
		am.cfg.Secret(),
		am.cfg.PassPhrase(),
		am.cfg.Key(),
		am.cfg.Version(),
	)
}
