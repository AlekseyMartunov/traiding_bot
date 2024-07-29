package httpclient

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	kucoinentity "tradingbot/internal/entity"
)

// marketOrderJSON helper dto struct.
type marketOrderJSON struct {
	// This field is returned when order information is obtained.
	// You can use clientOid to tag your orders.
	ClientOrderID string `json:"clientOid"`

	// buy or sell
	Side kucoinentity.Side `json:"side"`

	// e.g. ETH-BTC
	Symbol string `json:"symbol"`

	// limit or market (default is limit)
	OrderType string `json:"type"`

	Funds float64 `json:"funds"`
}

func (hc *HTTPClient) PlaceMarketOrder(order *kucoinentity.MarketOrder) error {
	body := marketOrderJSON{
		ClientOrderID: order.ClientOrderID,
		Side:          order.Side,
		Symbol:        order.Pair,
		OrderType:     "market",
		Funds:         order.Funds,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return hc.logAndReturnWrappedErr("marshal market order err", err)
	}

	headers := hc.secretHeaders(
		http.MethodPost,
		orderEndpoint,
		string(b),
		hc.cfg.Secret(),
		hc.cfg.PassPhrase(),
		hc.cfg.Key(),
		hc.cfg.Version(),
		time.Now(),
	)

	resp, err := hc.client.R().
		SetBody(b).
		SetHeaders(headers).
		Post(strings.Join([]string{hc.cfg.BaseEndpoint(), orderEndpoint}, ""))

	if err != nil {
		return hc.logAndReturnWrappedErr("market order request err", err)
	}

	_, err = hc.handleResponse(resp)
	if err != nil {
		hc.logAndReturnWrappedErr("handle market order response err", err)
	}

	return nil
}
