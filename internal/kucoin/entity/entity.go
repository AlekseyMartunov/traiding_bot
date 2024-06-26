// Package kucoinentity contains all main entity necessary to work kucoin bots.
package kucoinentity

import "time"

// Side can only have two meanings: sell or buy.
type Side string

const (
	Sell Side = "sell"
	Buy  Side = "buy"
)

// Ticker is used to obtain information about the currency and price.
type Ticker struct {
	Pair  string
	Price float64
}

// MlResult returns from the ьд service and indicates its decisions
// status true -> buy, status false -> sell
type MlResult struct {
	Pair   string
	Status bool
}

// MarketOrder is used to place an order on the kucoin exchange.
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

	Time time.Time
}

// OrderDetailInfo contains all info about the completed order.
type OrderDetailInfo struct {
	Id            string
	Symbol        string
	OpType        string
	Type          string
	Side          string
	Price         string
	Size          string
	Funds         string
	DealFunds     string
	DealSize      string
	Fee           string
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
	BaseMinSize     string
	QuoteMinSize    string
	BaseMaxSize     string
	QuoteMaxSize    string
	BaseIncrement   string
	QuoteIncrement  string
	PriceIncrement  string
	PriceLimitRate  string
	MinFunds        string
	IsMarginEnabled bool
	EnableTrading   bool
}
