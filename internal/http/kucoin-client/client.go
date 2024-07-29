package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type headerFunc func(method, url, body, secret, passPhrase, key, version string, now time.Time) map[string]string

const (
	orderEndpoint      = "/api/v1/orders"
	accountEndpoint    = "/api/v1/accounts"
	symbolListEndpoint = "/api/v2/symbols"
	candlesEndpoint    = "/api/v1/market/candles"
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
	Error(msg string, args ...any)
}

type rawMessage struct {
	Code string          `json:"code"`
	Data json.RawMessage `json:"data"`
}

type HTTPClient struct {
	log           logger
	cfg           config
	client        *resty.Client
	secretHeaders headerFunc
}

func New(l logger, c config, f headerFunc) *HTTPClient {
	return &HTTPClient{
		log:           l,
		cfg:           c,
		client:        resty.New(),
		secretHeaders: f,
	}
}

func (hc *HTTPClient) handleResponse(response *resty.Response) ([]byte, error) {
	if response.StatusCode() != http.StatusOK {
		return nil, errors.New("response status code not equal 200")
	}

	raw := &rawMessage{}
	if err := json.Unmarshal(response.Body(), raw); err != nil {
		return nil, errors.New(fmt.Sprintf("unmarsahl json err: %s", err.Error()))
	}

	hc.log.Debug(fmt.Sprintf("recive a message: %s", toStringFromJson(raw)))

	if raw.Code != successfulCode {
		hc.log.Error(fmt.Sprintf("kucoin exchange err: %s", toStringFromJson(raw)))
		return nil, errors.New(fmt.Sprintf("kucoin exchange err: %s", toStringFromJson(raw)))
	}

	return raw.Data, nil
}

func (hc *HTTPClient) logAndReturnWrappedErr(msg string, e error) error {
	newErr := fmt.Errorf("%s: %w", msg, e)
	hc.log.Error(newErr.Error())
	return newErr
}

func toStringFromJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
