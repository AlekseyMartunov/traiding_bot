// Package kucoinentity contains all main entity necessary to work kucoin-client bots.
package kucoinentity

import "time"

// Side can only have two meanings: sell or buy.
type Side string

const (
	Sell Side = "sell"
	Buy  Side = "buy"
)

// MarketOrder is used to place an order on the kucoin-client exchange.
type MarketOrder struct {
	// order id returned by server, ClientOrderID is different from the OrderID
	OrderID string

	// client side uuid for order
	ClientOrderID string

	// sell or buy
	Side Side

	// Funds field refers to the funds for the priced asset (the asset name written latter) of the trading pair.
	// Example: "funds: 25,  pair: BTC-USD, side: buy" means that we want buy BTC by 25 USDT
	Funds float64

	// BTC-USDT
	Pair string

	// Remark, length cannot exceed 50 characters (ASCII)
	Remark string

	Time time.Time
}

type StopOrderTrigger string

// loss : Triggers when the last trade price changes to a value at or below the stopPrice.
// entry : Triggers when the last trade price changes to a value at or above the stopPrice.
const (
	loss  StopOrderTrigger = "loss"
	entry StopOrderTrigger = "entry"
)

type StopOrder struct {
	// order id returned by server, ClientOrderID is different from the OrderID
	OrderID string

	// client side uuid for order
	ClientOrderID string

	// sell or buy
	Side Side

	// BTC-USDT
	Pair string

	// Order placement remarks, length cannot exceed 100 characters (UTF-8)
	Remark string

	// Either loss or entry, the default is loss. Requires stopPrice to be defined.
	Stop StopOrderTrigger

	// trigger price.
	StopPrice float64

	// Specify quantity for currency
	Size float64

	// Specify price for currency
	Price float64

	Time time.Time
}

// OrderDetailInfo contains all info about the completed order.
type OrderDetailInfo struct {
	Id            string
	Symbol        string
	OpType        string
	Type          string
	Side          string
	Price         float64
	Size          float64
	Funds         float64
	DealFunds     float64
	DealSize      float64
	Fee           float64
	FeeCurrency   string
	Stp           string
	Stop          string
	StopTriggered bool
	StopPrice     string
	TimeInForce   string
	PostOnly      bool
	Hidden        bool
	Iceberg       bool
	VisibleSize   string
	CancelAfter   int
	Channel       string
	ClientOid     string
	Remark        interface{}
	Tags          interface{}
	IsActive      bool
	CancelExist   bool
	CreatedAt     time.Time
	TradeType     string
}

// AccountInfo contains info about currency accounts.
type AccountInfo struct {
	ID        string
	Currency  string
	Type      string
	Balance   float64
	Available float64
	Holds     float64
}

// CurrencyConfig contains info on purchase and sales volumes.
type CurrencyConfig struct {
	Symbol          string
	Name            string
	BaseCurrency    string
	QuoteCurrency   string
	FeeCurrency     string
	Market          string
	BaseMinSize     float64
	QuoteMinSize    float64
	BaseMaxSize     float64
	QuoteMaxSize    float64
	BaseIncrement   float64
	QuoteIncrement  float64
	PriceIncrement  float64
	PriceLimitRate  float64
	MinFunds        float64
	IsMarginEnabled bool
	EnableTrading   bool
}

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
