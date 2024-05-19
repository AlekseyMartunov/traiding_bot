package kucoinorders

type marketOrderForRequest struct {
	// This field is returned when order information is obtained.
	// You can use clientOid to tag your orders.
	OrderUUID string `json:"client_oid"`

	// buy or sell
	Side string `json:"side"`

	// e.g. ETH-BTC
	Symbol string `json:"symbol"`

	// limit or market (default is limit)
	OrderType string `json:"type"`

	Size string `json:"size"`
}

type responseOrder struct {
	OrderID string `json:"orderId"`
}
