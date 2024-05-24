package entity

import "time"

type Side string

const (
	Sell Side = "sell"
	Buy  Side = "buy"
)

type Ticker struct {
	Pair  string
	Price string
}

type MarketOrder struct {
	// order id returned by server, ClientOrderID is different from the OrderID
	OrderID string

	// client side uuid for order
	ClientOrderID string

	Side  Side
	Funds float64
	Size  float64
	Pair  string
	Time  time.Time
}
