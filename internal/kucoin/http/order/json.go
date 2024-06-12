package kucoinorders

import (
	"strconv"
	"time"
	"tradingbot/internal/kucoin/entity"
)

// marketOrderJSON helper dto struct.
type marketOrderJSON struct {
	// This field is returned when order information is obtained.
	// You can use clientOid to tag your orders.
	ClientOrderID string `json:"clientOid"`

	// buy or sell
	Side kucoinentity.Side `json:"side"`

	// e.g. ETH-BTC
	Symbol string `json:"symbol"`

	// limit or market (default is limit)
	OrderType string `json:"type"`

	Funds string `json:"funds"`
}

// responseOrderJSON helper dto struct.
type responseOrderJSON struct {
	Code string `json:"code"`
	Data struct {
		OrderId string `json:"orderId"`
	} `json:"data"`
}

// orderDetailInfoJSON helper dto struct.
type orderDetailInfoJSON struct {
	Code string `json:"code"`
	Data struct {
		Id            string      `json:"id"`
		Symbol        string      `json:"symbol"`
		OpType        string      `json:"opType"`
		Type          string      `json:"type"`
		Side          string      `json:"side"`
		Price         string      `json:"price"`
		Size          string      `json:"size"`
		Funds         string      `json:"funds"`
		DealFunds     string      `json:"dealFunds"`
		DealSize      string      `json:"dealSize"`
		Fee           string      `json:"fee"`
		FeeCurrency   string      `json:"feeCurrency"`
		Stp           string      `json:"stp"`
		Stop          string      `json:"stop"`
		StopTriggered bool        `json:"stopTriggered"`
		StopPrice     string      `json:"stopPrice"`
		TimeInForce   string      `json:"timeInForce"`
		PostOnly      bool        `json:"postOnly"`
		Hidden        bool        `json:"hidden"`
		Iceberg       bool        `json:"iceberg"`
		VisibleSize   string      `json:"visibleSize"`
		CancelAfter   int         `json:"cancelAfter"`
		Channel       string      `json:"channel"`
		ClientOid     string      `json:"clientOid"`
		Remark        interface{} `json:"remark"`
		Tags          interface{} `json:"tags"`
		IsActive      bool        `json:"isActive"`
		CancelExist   bool        `json:"cancelExist"`
		CreatedAt     int64       `json:"createdAt"`
		TradeType     string      `json:"tradeType"`
	} `json:"data"`
}

func (o *orderDetailInfoJSON) toBaseEntity() (*kucoinentity.OrderDetailInfo, error) {
	var result kucoinentity.OrderDetailInfo

	orderTimeIntFormat, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		return nil, err
	}

	result.Id = o.Data.Id
	result.Symbol = o.Data.Symbol
	result.OpType = o.Data.OpType
	result.Type = o.Data.Type
	result.Side = o.Data.Side
	result.Price = o.Data.Price
	result.Size = o.Data.Size
	result.Funds = o.Data.Funds
	result.DealFunds = o.Data.DealFunds
	result.DealSize = o.Data.DealSize
	result.Fee = o.Data.Fee
	result.FeeCurrency = o.Data.FeeCurrency
	result.Stp = o.Data.Stp
	result.Stop = o.Data.Stop
	result.StopTriggered = o.Data.StopTriggered
	result.StopPrice = o.Data.StopPrice
	result.TimeInForce = o.Data.TimeInForce
	result.PostOnly = o.Data.PostOnly
	result.Hidden = o.Data.Hidden
	result.Iceberg = o.Data.Iceberg
	result.VisibleSize = o.Data.VisibleSize
	result.CancelAfter = o.Data.CancelAfter
	result.Channel = o.Data.Channel
	result.ClientOid = o.Data.ClientOid
	result.Remark = o.Data.Remark
	result.Tags = o.Data.Tags
	result.IsActive = o.Data.IsActive
	result.CancelExist = o.Data.CancelExist
	result.CreatedAt = time.Unix(orderTimeIntFormat, 0)
	result.TradeType = o.Data.TradeType

	return &result, nil
}
