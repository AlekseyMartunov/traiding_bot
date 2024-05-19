package entity

import "time"

type Ticker struct {
	Pair  string
	Price string
}

type MarketOrder struct {
	OrderID   string
	Price     float64
	Pair      string
	OpenTime  time.Time
	CloseTime time.Time
}
