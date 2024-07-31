package httpclient

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	kucoinentity "tradingbot/internal/entity"
)

type stopOrderJSON struct {
	ClientOrderID string                        `json:"clientOid"`
	Side          kucoinentity.Side             `json:"side"`
	Symbol        string                        `json:"symbol"`
	Type          string                        `json:"type"`
	Remark        string                        `json:"remark"`
	Stop          kucoinentity.StopOrderTrigger `json:"stop"`
	StopPrice     float64                       `json:"stopPrice"`

	// Specify price for currency
	Price float64 `json:"price"`

	// Specify quantity for currency
	Size float64 `json:"size"`
}

func (hc *HTTPClient) PlaceStopOrder(o *kucoinentity.StopOrder) error {
	body := stopOrderJSON{
		ClientOrderID: o.ClientOrderID,
		Side:          o.Side,
		Symbol:        o.Pair,
		Type:          "limit",
		Remark:        "",
		Stop:          o.Stop,
		StopPrice:     o.StopPrice,
		Price:         o.Price,
		Size:          o.Size,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return hc.logAndReturnWrappedErr("marshal stop order err", err)
	}

	headers := hc.secretHeaders(
		http.MethodPost,
		stopOrderEndpoint,
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
		Post(strings.Join([]string{hc.cfg.BaseEndpoint(), stopOrderEndpoint}, ""))

	if err != nil {
		return hc.logAndReturnWrappedErr("stop order request err", err)
	}

	b, err = hc.handleResponse(resp)
	if err != nil {
		hc.logAndReturnWrappedErr("handle stop order response err", err)
	}

	j := &orderIDJSON{}

	if err = json.Unmarshal(b, j); err != nil {
		return hc.logAndReturnWrappedErr("unmarshal stop order err", err)
	}

	o.OrderID = j.orderId

	return nil
}
