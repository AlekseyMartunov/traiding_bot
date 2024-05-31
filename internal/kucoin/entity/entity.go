package kucoinentity

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

type AccountInfo struct {
	ID        string
	Currency  string
	Type      string
	Balance   float64
	Available float64
	Holds     float64
}

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
	CreatedAt     int64
	TradeType     string
}
