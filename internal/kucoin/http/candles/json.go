package kucoincandles

import (
	"strconv"
	"time"

	kucoinentity "tradingbot/internal/kucoin/entity"
)

type candlesJSON struct {
	Code string     `json:"code"`
	Data [][]string `json:"data"`
}

func (cj *candlesJSON) toBaseEntity() ([]*kucoinentity.Candle, error) {
	res := make([]*kucoinentity.Candle, 0, len(cj.Data))

	for _, obj := range cj.Data {
		e, err := parseCandle(obj)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, nil
}

func parseCandle(s []string) (*kucoinentity.Candle, error) {
	i, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return nil, err
	}
	time := time.Unix(i, 0)

	open, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return nil, err
	}

	close, err := strconv.ParseFloat(s[2], 64)
	if err != nil {
		return nil, err
	}

	high, err := strconv.ParseFloat(s[3], 64)
	if err != nil {
		return nil, err
	}

	low, err := strconv.ParseFloat(s[4], 64)
	if err != nil {
		return nil, err
	}

	volume, err := strconv.ParseFloat(s[5], 64)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(s[6], 64)
	if err != nil {
		return nil, err
	}

	var c = kucoinentity.Candle{}

	c.Open = open
	c.Close = close
	c.High = high
	c.Low = low
	c.Turnover = amount
	c.Volume = volume
	c.Time = time

	return &c, nil
}
