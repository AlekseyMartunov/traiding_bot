package httpclient

import (
	"encoding/json"
	"strings"
	kucoinentity "tradingbot/internal/entity"
)

func (hc *HTTPClient) Currency(pair string) (*kucoinentity.CurrencyConfig, error) {
	url := strings.Join([]string{hc.cfg.BaseEndpoint(), symbolListEndpoint, "/", pair}, "")

	response, err := hc.client.R().Get(url)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("create currency request err", err)
	}

	b, err := hc.handleResponse(response)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("handle currency response err", err)
	}

	res := &kucoinentity.CurrencyConfig{}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal currency response err", err)
	}

	return res, nil

}
