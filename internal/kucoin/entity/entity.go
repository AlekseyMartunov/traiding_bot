// Package kucoinentity contains all main entity necessary to work kucoin bots.
package kucoinentity

// Ticker is used to obtain information about the currency and price.
type Ticker struct {
	Pair  string
	Price float64
}
