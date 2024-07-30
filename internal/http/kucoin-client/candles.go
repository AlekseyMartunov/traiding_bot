package httpclient

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	entity "tradingbot/internal/entity"
)

type candlesJSON struct {
	candles [][7]string
}

func (cj *candlesJSON) toBaseEntity() ([]*entity.Candle, error) {
	result := make([]*entity.Candle, 0, len(cj.candles))
	for _, c := range cj.candles {
		baseEntity, err := cj.tryToConvert(c)
		if err != nil {
			return nil, err
		}
		result = append(result, baseEntity)
	}
	return result, nil
}

func (cj *candlesJSON) tryToConvert(candela [7]string) (*entity.Candle, error) {
	seconds, err := strconv.ParseInt(candela[0], 10, 64)
	if err != nil {
		return nil, err
	}

	t := time.Unix(seconds, 0)

	open, err := strconv.ParseFloat(candela[1], 10)
	if err != nil {
		return nil, err
	}

	closePrice, err := strconv.ParseFloat(candela[2], 10)
	if err != nil {
		return nil, err
	}

	high, err := strconv.ParseFloat(candela[3], 10)
	if err != nil {
		return nil, err
	}

	low, err := strconv.ParseFloat(candela[4], 10)
	if err != nil {
		return nil, err
	}

	volume, err := strconv.ParseFloat(candela[5], 10)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(candela[5], 10)
	if err != nil {
		return nil, err
	}

	return &entity.Candle{
		Time:     t,
		Open:     open,
		Close:    closePrice,
		High:     high,
		Low:      low,
		Volume:   volume,
		Turnover: amount,
	}, nil

}

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

	jsonArr := candlesJSON{}

	err = json.Unmarshal(b, &jsonArr.candles)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal candles err", err)
	}

	base, err := jsonArr.toBaseEntity()
	if err != nil {
		hc.logAndReturnWrappedErr("candles parsing err", err)
	}

	return base, nil

}
