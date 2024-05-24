package kucoinorders

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"strings"
	"tradingbot/internal/kucoin/entity"
)

const (
	testEndpoint = "/api/v1/orders/test"
	endpoint     = "/api/v1/orders"
	baseEndpoint = "https://api.kucoin.com"
)

var (
	StatusCodeIsNot200 = errors.New("response status code not equal 200")
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

type KucoinMarketPlacer struct {
	log      logger
	cfg      config
	client   *resty.Client
	endpoint string
}

func NewKucoinMarketPlacer(l logger, c config) *KucoinMarketPlacer {
	return &KucoinMarketPlacer{
		log:      l,
		cfg:      c,
		client:   resty.New(),
		endpoint: endpoint,
	}
}

func (p *KucoinMarketPlacer) PlaceMarketOrder(order *entity.MarketOrder) error {
	body := marketOrderForRequest{
		ClientOrderID: order.ClientOrderID,
		Side:          order.Side,
		Symbol:        order.Pair,
		OrderType:     "market",
		Funds:         strconv.FormatFloat(order.Funds, 'f', 6, 64),
		//Size:          strconv.FormatFloat(order.Size, 'f', 6, 64),
	}

	b, err := json.Marshal(body)
	if err != nil {
		p.log.Error(err.Error())
		return err
	}
	fmt.Println(string(b))

	headers := p.createHeaders(http.MethodPost, p.endpoint, string(b))

	response, err := p.client.R().
		SetBody(b).
		SetHeaders(headers).
		Post(strings.Join([]string{baseEndpoint, p.endpoint}, ""))

	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	if response.StatusCode() != 200 {
		p.log.Info(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return err
	}

	respOrder := responseOrder{}

	err = json.Unmarshal(response.Body(), &respOrder)
	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	order.OrderID = respOrder.Data.OrderId
	fmt.Println(response.String())

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

func (p *KucoinMarketPlacer) Test() error {
	e := "/api/v1/accounts"
	headers := p.createHeaders(http.MethodGet, e, "")
	resp, err := p.client.R().SetHeaders(headers).Get(strings.Join([]string{baseEndpoint, e}, ""))
	if err != nil {
		fmt.Println("ERR:", err.Error())
	}

	fmt.Println(resp.String())

	return err

}
