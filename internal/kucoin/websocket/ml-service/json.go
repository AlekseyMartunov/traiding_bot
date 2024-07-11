package ml

import (
	kucoinentity "tradingbot/internal/kucoin/entity"
)

type tickerJSON struct {
	Pair  string  `json:"pair"`
	Price float64 `json:"price"`
}

func (t *tickerJSON) fromEntity(base *kucoinentity.Ticker) {
	t.Pair = base.Pair
	t.Price = base.Price
}

func (t *tickerJSON) toEntity() *kucoinentity.Ticker {
	base := kucoinentity.Ticker{
		Pair:  t.Pair,
		Price: t.Price,
	}

	return &base
}
