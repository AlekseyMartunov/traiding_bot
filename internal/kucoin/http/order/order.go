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

	//Size string `json:"size,omitempty"`
}

type responseOrder struct {
	Code string `json:"code"`
	Data struct {
		OrderId string `json:"orderId"`
	} `json:"data"`
}
