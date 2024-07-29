package httpclient

import (
	"fmt"
	"strings"
	"time"

	entity "tradingbot/internal/entity"
)

func (hc *HTTPClient) Candles(
	pair string,
	timeFrame entity.CandlePeriod, from, to time.Time) ([]*entity.Candle, error) {

	url := strings.Join([]string{hc.cfg.BaseEndpoint(), candlesEndpoint}, "")

	resp, err := hc.client.R().
		SetQueryParam("type", string(timeFrame)).
		SetQueryParam("symbol", pair).
		SetQueryParam("startAt", fmt.Sprint(from.Unix())).
		SetQueryParam("endAt", fmt.Sprint(to.Unix())).
		Get(url)

	if err != nil {
		return nil, hc.logAndReturnWrappedErr("create candles request err", err)
	}

	b, err := hc.handleResponse(resp)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("handles candles response err", err)
	}

	hc.log.Debug(string(b))

	return nil, nil

}
