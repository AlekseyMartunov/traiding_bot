package httpclient

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	kucoinentity "tradingbot/internal/entity"
)

type orderDetailJSON struct {
	Id            string            `json:"id"`
	Symbol        string            `json:"symbol"`
	OpType        string            `json:"opType"`
	Type          string            `json:"type"`
	Side          kucoinentity.Side `json:"side"`
	Price         string            `json:"price"`
	Size          string            `json:"size"`
	Funds         string            `json:"funds"`
	DealFunds     string            `json:"dealFunds"`
	DealSize      string            `json:"dealSize"`
	Fee           string            `json:"fee"`
	FeeCurrency   string            `json:"feeCurrency"`
	Stp           string            `json:"stp"`
	Stop          string            `json:"stop"`
	StopTriggered bool              `json:"stopTriggered"`
	StopPrice     string            `json:"stopPrice"`
	TimeInForce   string            `json:"timeInForce"`
	PostOnly      bool              `json:"postOnly"`
	Hidden        bool              `json:"hidden"`
	Iceberg       bool              `json:"iceberg"`
	VisibleSize   string            `json:"visibleSize"`
	CancelAfter   int               `json:"cancelAfter"`
	Channel       string            `json:"channel"`
	ClientOid     string            `json:"clientOid"`
	Remark        string            `json:"remark"`
	Tags          string            `json:"tags"`
	IsActive      bool              `json:"isActive"`
	CancelExist   bool              `json:"cancelExist"`
	CreatedAt     int64             `json:"createdAt"`
	TradeType     string            `json:"tradeType"`
}

func (o *orderDetailJSON) toBaseEntity() (*kucoinentity.OrderDetailInfo, error) {
	var result kucoinentity.OrderDetailInfo

	Price, err := strconv.ParseFloat(o.Price, 10)
	if err != nil {
		return nil, err
	}

	Size, err := strconv.ParseFloat(o.Size, 10)
	if err != nil {
		return nil, err
	}

	Funds, err := strconv.ParseFloat(o.Funds, 10)
	if err != nil {
		return nil, err
	}

	DealFunds, err := strconv.ParseFloat(o.DealFunds, 10)
	if err != nil {
		return nil, err
	}

	DealSize, err := strconv.ParseFloat(o.DealSize, 10)
	if err != nil {
		return nil, err
	}

	Fee, err := strconv.ParseFloat(o.Fee, 10)
	if err != nil {
		return nil, err
	}

	result.Id = o.Id
	result.Symbol = o.Symbol
	result.OpType = o.OpType
	result.Type = o.Type
	result.Side = o.Side
	result.Price = Price
	result.Size = Size
	result.Funds = Funds
	result.DealFunds = DealFunds
	result.DealSize = DealSize
	result.Fee = Fee
	result.FeeCurrency = o.FeeCurrency
	result.Stp = o.Stp
	result.Stop = o.Stop
	result.StopTriggered = o.StopTriggered
	result.StopPrice = o.StopPrice
	result.TimeInForce = o.TimeInForce
	result.PostOnly = o.PostOnly
	result.Hidden = o.Hidden
	result.Iceberg = o.Iceberg
	result.VisibleSize = o.VisibleSize
	result.CancelAfter = o.CancelAfter
	result.Channel = o.Channel
	result.ClientOid = o.ClientOid
	result.Remark = o.Remark
	result.Tags = o.Tags
	result.IsActive = o.IsActive
	result.CancelExist = o.CancelExist
	result.CreatedAt = time.Unix(o.CreatedAt, 0)
	result.TradeType = o.TradeType

	return &result, nil
}

func (hc *HTTPClient) OrderDetail(orderID string) (*kucoinentity.OrderDetailInfo, error) {
	url := strings.Join([]string{orderEndpoint, "/", orderID}, "")

	headers := hc.secretHeaders(
		http.MethodGet,
		url,
		"",
		hc.cfg.Secret(),
		hc.cfg.PassPhrase(),
		hc.cfg.Key(),
		hc.cfg.Version(),
		time.Now(),
	)

	response, err := hc.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{hc.cfg.BaseEndpoint(), url}, ""))

	if err != nil {
		return nil, hc.logAndReturnWrappedErr("order detail request err", err)
	}

	b, err := hc.handleResponse(response)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("handle order detail response err", err)
	}

	j := &orderDetailJSON{}

	if err = json.Unmarshal(b, j); err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal order detail err", err)
	}

	base, err := j.toBaseEntity()
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("order detail parsing err", err)
	}

	return base, nil

}
