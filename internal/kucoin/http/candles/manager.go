package kucoincandles

import "github.com/go-resty/resty/v2"

const (
	endpoint = "/api/v1/market/candles"
)

type config interface {
	GetBaseEndpoint() string
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Event(msg string, args ...any)
	Error(msg string, args ...any)
}

type CandlesManager struct {
	log    logger
	cfg    config
	client *resty.Client
}

func New(l logger, c config) *CandlesManager {
	return &CandlesManager{
		log:    l,
		cfg:    c,
		client: resty.New(),
	}
}
