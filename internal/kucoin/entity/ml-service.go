package kucoinentity

import "encoding/json"

type MlServiceRawMessage struct {
	// Can be: "event", "historical"
	Type string `json:"type"`

	Data json.RawMessage `json:"data"`
}

type MlServiceResponseEvent struct {
	Currency string `json:"currency"`
	Side     Side   `json:"side"`
}
