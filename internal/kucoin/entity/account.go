package kucoinentity

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
