package kucoinaccount

import (
	"strconv"
	kucoinentity "tradingbot/internal/kucoin/entity"
)

type accountInfoJSON struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

func (a *accountInfoJSON) toBaseEntity() (*kucoinentity.AccountInfo, error) {
	var result kucoinentity.AccountInfo

	balance, err := strconv.ParseFloat(a.Balance, 64)
	if err != nil {
		return nil, err
	}

	available, err := strconv.ParseFloat(a.Available, 64)
	if err != nil {
		return nil, err
	}

	holds, err := strconv.ParseFloat(a.Holds, 64)
	if err != nil {
		return nil, err
	}

	result.ID = a.Id
	result.Currency = a.Currency
	result.Type = a.Type
	result.Balance = balance
	result.Available = available
	result.Holds = holds

	return &result, nil
}
