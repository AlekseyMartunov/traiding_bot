package ml

import kucoinentity "tradingbot/internal/kucoin/entity"

type tickerDTO struct {
	Pair  string  `json:"pair"`
	Price float64 `json:"price"`
}

func (t *tickerDTO) fromEntity(base *kucoinentity.Ticker) {
	t.Pair = base.Pair
	t.Price = base.Price
}

func (t *tickerDTO) toEntity() *kucoinentity.Ticker {
	base := kucoinentity.Ticker{
		Pair:  t.Pair,
		Price: t.Price,
	}

	return &base
}

type mlResultDTO struct {
	Pair   string `json:"pair"`
	Status bool   `json:"status"`
}

func (m *mlResultDTO) toEntity() *kucoinentity.MlResult {
	base := kucoinentity.MlResult{
		Pair:   m.Pair,
		Status: m.Status,
	}

	return &base
}
