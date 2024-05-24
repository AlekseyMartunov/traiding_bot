package kucoinorders

import "tradingbot/internal/kucoin/entity"

type marketOrderForRequest struct {
	// This field is returned when order information is obtained.
	// You can use clientOid to tag your orders.
	ClientOrderID string `json:"clientOid"`

	// buy or sell
	Side entity.Side `json:"side"`

	// e.g. ETH-BTC
	Symbol string `json:"symbol"`

	// limit or market (default is limit)
	OrderType string `json:"type"`

	Funds string `json:"funds"`
}

type responseOrder struct {
	Code string `json:"code"`
	Data struct {
		OrderId string `json:"orderId"`
	} `json:"data"`
}

type currencyConfig struct {
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

type orderDetailInfo struct {
	Code string `json:"code"`
	Data struct {
		Id            string      `json:"id"`
		Symbol        string      `json:"symbol"`
		OpType        string      `json:"opType"`
		Type          string      `json:"type"`
		Side          string      `json:"side"`
		Price         string      `json:"price"`
		Size          string      `json:"size"`
		Funds         string      `json:"funds"`
		DealFunds     string      `json:"dealFunds"`
		DealSize      string      `json:"dealSize"`
		Fee           string      `json:"fee"`
		FeeCurrency   string      `json:"feeCurrency"`
		Stp           string      `json:"stp"`
		Stop          string      `json:"stop"`
		StopTriggered bool        `json:"stopTriggered"`
		StopPrice     string      `json:"stopPrice"`
		TimeInForce   string      `json:"timeInForce"`
		PostOnly      bool        `json:"postOnly"`
		Hidden        bool        `json:"hidden"`
		Iceberg       bool        `json:"iceberg"`
		VisibleSize   string      `json:"visibleSize"`
		CancelAfter   int         `json:"cancelAfter"`
		Channel       string      `json:"channel"`
		ClientOid     string      `json:"clientOid"`
		Remark        interface{} `json:"remark"`
		Tags          interface{} `json:"tags"`
		IsActive      bool        `json:"isActive"`
		CancelExist   bool        `json:"cancelExist"`
		CreatedAt     int64       `json:"createdAt"`
		TradeType     string      `json:"tradeType"`
	} `json:"data"`
}
