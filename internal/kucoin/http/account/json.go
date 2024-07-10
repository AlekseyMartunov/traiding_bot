package kucoinaccount

import (
	"strconv"
	kucoinentity "tradingbot/internal/kucoin/entity"
)

// accountInfoJSON  helper dto struct.
type accountInfoJSON struct {
	Code string `json:"code"`
	Data []struct {
		Id        string `json:"id"`
		Currency  string `json:"currency"`
		Type      string `json:"type"`
		Balance   string `json:"balance"`
		Available string `json:"available"`
		Holds     string `json:"holds"`
	} `json:"data"`
}

func (a *accountInfoJSON) toBaseEntity() ([]*kucoinentity.AccountInfo, error) {
	result := make([]*kucoinentity.AccountInfo, 0, len(a.Data))

	for _, obj := range a.Data {
		balance, err := strconv.ParseFloat(obj.Balance, 64)
		if err != nil {
			return nil, err
		}

		available, err := strconv.ParseFloat(obj.Available, 64)
		if err != nil {
			return nil, err
		}

		holds, err := strconv.ParseFloat(obj.Holds, 64)
		if err != nil {
			return nil, err
		}

		e := kucoinentity.AccountInfo{}

		e.ID = obj.Id
		e.Currency = obj.Currency
		e.Type = obj.Type
		e.Balance = balance
		e.Available = available
		e.Holds = holds

		result = append(result, &e)
	}

	return result, nil
}

// currencyConfigJSON helper dto struct.
type currencyConfigJSON struct {
	Code string `json:"code"`
	Data struct {
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
	} `json:"data"`
}

func (c *currencyConfigJSON) toBaseEntity() *kucoinentity.CurrencyConfig {
	var result kucoinentity.CurrencyConfig

	result.Symbol = c.Data.Symbol
	result.Name = c.Data.Name
	result.BaseCurrency = c.Data.BaseCurrency
	result.QuoteCurrency = c.Data.QuoteCurrency
	result.FeeCurrency = c.Data.FeeCurrency
	result.Market = c.Data.Market
	result.BaseMinSize = c.Data.BaseMinSize
	result.QuoteMinSize = c.Data.QuoteMinSize
	result.BaseMaxSize = c.Data.BaseMaxSize
	result.QuoteMaxSize = c.Data.QuoteMaxSize
	result.BaseIncrement = c.Data.BaseIncrement
	result.QuoteIncrement = c.Data.QuoteIncrement
	result.PriceIncrement = c.Data.PriceIncrement
	result.PriceLimitRate = c.Data.PriceLimitRate
	result.MinFunds = c.Data.MinFunds
	result.IsMarginEnabled = c.Data.IsMarginEnabled
	result.EnableTrading = c.Data.EnableTrading

	return &result
}
