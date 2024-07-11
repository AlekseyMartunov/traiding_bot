package kucoinentity

import "time"

type CandlePeriod string

const (
	Min1   CandlePeriod = "1min"
	Min3   CandlePeriod = "3min"
	Min5   CandlePeriod = "5min"
	Min15  CandlePeriod = "15min"
	Min30  CandlePeriod = "30min"
	Hour1  CandlePeriod = "1hour"
	Hour2  CandlePeriod = "2hour"
	Hour4  CandlePeriod = "4hour"
	Hour6  CandlePeriod = "6hour"
	Hour8  CandlePeriod = "8hour"
	Hour12 CandlePeriod = "12hour"
	Day1   CandlePeriod = "1day"
	Week1  CandlePeriod = "1week"
	Mouth1 CandlePeriod = "1month"
)

type Candle struct {
	// Start time of the candle cycle
	Time time.Time

	Open  float64
	Close float64
	High  float64
	Low   float64

	// Transaction volume(One-sided transaction volume)
	Volume float64

	// Transaction amount(One-sided transaction amount)
	Turnover float64
}
