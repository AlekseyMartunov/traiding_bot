package httpclient

import (
	"encoding/json"
	"strconv"
	"strings"

	kucoinentity "tradingbot/internal/entity"
)

type currencyJSON struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	FeeCurrency     string `json:"feeCurrency"`
	Market          string `json:"market"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteMinSize    string `json:"quoteMinSize"`
	BaseMaxSize     string `json:"baseMaxSize"`
	QuoteMaxSize    string `json:"quoteMaxSize"`
	BaseIncrement   string `json:"baseIncrement"`
	QuoteIncrement  string `json:"quoteIncrement"`
	PriceIncrement  string `json:"priceIncrement"`
	PriceLimitRate  string `json:"priceLimitRate"`
	MinFunds        string `json:"minFunds"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	EnableTrading   bool   `json:"enableTrading"`
}

func (cj *currencyJSON) toBaseEntity() (*kucoinentity.CurrencyConfig, error) {
	BaseMinSize, err := strconv.ParseFloat(cj.BaseMinSize, 10)
	if err != nil {
		return nil, err
	}

	QuoteMinSize, err := strconv.ParseFloat(cj.QuoteMinSize, 10)
	if err != nil {
		return nil, err
	}

	BaseMaxSize, err := strconv.ParseFloat(cj.BaseMaxSize, 10)
	if err != nil {
		return nil, err
	}

	QuoteMaxSize, err := strconv.ParseFloat(cj.QuoteMaxSize, 10)
	if err != nil {
		return nil, err
	}

	BaseIncrement, err := strconv.ParseFloat(cj.BaseIncrement, 10)
	if err != nil {
		return nil, err
	}

	QuoteIncrement, err := strconv.ParseFloat(cj.QuoteIncrement, 10)
	if err != nil {
		return nil, err
	}

	PriceIncrement, err := strconv.ParseFloat(cj.PriceIncrement, 10)
	if err != nil {
		return nil, err
	}

	PriceLimitRate, err := strconv.ParseFloat(cj.PriceLimitRate, 10)
	if err != nil {
		return nil, err
	}

	MinFunds, err := strconv.ParseFloat(cj.MinFunds, 10)
	if err != nil {
		return nil, err
	}

	return &kucoinentity.CurrencyConfig{
		Symbol:          cj.Symbol,
		Name:            cj.Name,
		BaseCurrency:    cj.BaseCurrency,
		QuoteCurrency:   cj.QuoteCurrency,
		FeeCurrency:     cj.FeeCurrency,
		Market:          cj.Market,
		BaseMinSize:     BaseMinSize,
		QuoteMinSize:    QuoteMinSize,
		BaseMaxSize:     BaseMaxSize,
		QuoteMaxSize:    QuoteMaxSize,
		BaseIncrement:   BaseIncrement,
		QuoteIncrement:  QuoteIncrement,
		PriceIncrement:  PriceIncrement,
		PriceLimitRate:  PriceLimitRate,
		MinFunds:        MinFunds,
		IsMarginEnabled: false,
		EnableTrading:   false,
	}, nil
}

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

	j := &currencyJSON{}
	if err := json.Unmarshal(b, j); err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal currency response err", err)
	}

	base, err := j.toBaseEntity()
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("currency parsing err", err)
	}

	return base, nil

}
