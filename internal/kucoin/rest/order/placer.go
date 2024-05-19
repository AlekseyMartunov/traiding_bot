package kucoinorders

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"tradingbot/internal/kucoin/entity"
)

const (
	testEndpoint = "https://api.kucoin.com/api/v1/orders/test"
	endpoint     = "https://api.kucoin.com/api/v1/orders"
)

var (
	StatusCodeIsNot200 = errors.New("response status code not equal 200")
)

type config interface {
	Key() string
	Secret() string
	Version() string
	PassPhrase() string
	AccountId() string
}

type logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
}

type KucoinMarketPlacer struct {
	log    logger
	client *resty.Client
	cfg    config
}

func NewKucoinMarketPlacer(l logger, c config) *KucoinMarketPlacer {
	return &KucoinMarketPlacer{
		log:    l,
		cfg:    c,
		client: resty.New(),
	}
}

func (p *KucoinMarketPlacer) OpenOrder(order *entity.MarketOrder) error {
	body := marketOrderForRequest{
		OrderUUID: order.OrderID,
		Side:      "buy",
		Symbol:    order.Pair,
		OrderType: "market",
		Size:      strconv.FormatFloat(order.Price, 'f', -1, 64),
	}

	b, err := json.Marshal(body)
	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	headers := p.createHeaders("post", endpoint, string(b))

	response, err := p.client.R().
		SetBody(b).
		SetHeaders(headers).
		Post(endpoint)

	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	if response.StatusCode() != 200 {
		p.log.Info(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return err
	}

	return nil
}

func (p *KucoinMarketPlacer) CloseOrder(order *entity.MarketOrder) error {
	body := marketOrderForRequest{
		OrderUUID: order.OrderID,
		Side:      "sell",
		Symbol:    order.Pair,
		OrderType: "market",
		Size:      strconv.FormatFloat(order.Price, 'f', -1, 64),
	}

	b, err := json.Marshal(body)
	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	headers := p.createHeaders("post", endpoint, string(b))

	response, err := p.client.R().
		SetBody(b).
		SetHeaders(headers).
		Post(endpoint)

	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	if response.StatusCode() != 200 {
		p.log.Info(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return err
	}
	return nil
}

func (p *KucoinMarketPlacer) createHeaders(method, url, body string) map[string]string {
	return createSecretsHeaders(
		method,
		url,
		body,
		p.cfg.Secret(),
		p.cfg.PassPhrase(),
		p.cfg.Key(),
		p.cfg.Version(),
	)
}
