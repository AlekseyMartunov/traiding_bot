package httpclient

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	kucoinentity "tradingbot/internal/entity"
)

type stopOrderDetailJSON struct {
	Id              string                        `json:"id"`
	Symbol          string                        `json:"symbol"`
	UserId          string                        `json:"userId"`
	Status          string                        `json:"status"`
	Type            string                        `json:"type"`
	Side            kucoinentity.Side             `json:"side"`
	Price           string                        `json:"price"`
	Size            string                        `json:"size"`
	Funds           string                        `json:"funds"`
	Stp             string                        `json:"stp"`
	TimeInForce     string                        `json:"timeInForce"`
	CancelAfter     int64                         `json:"cancelAfter"`
	PostOnly        bool                          `json:"postOnly"`
	Hidden          bool                          `json:"hidden"`
	Iceberg         bool                          `json:"iceberg"`
	VisibleSize     string                        `json:"visibleSize"`
	Channel         string                        `json:"channel"`
	ClientOid       string                        `json:"clientOid"`
	Remark          string                        `json:"remark"`
	Tags            string                        `json:"tags"`
	OrderTime       int64                         `json:"orderTime"`
	DomainId        string                        `json:"domainId"`
	TradeSource     string                        `json:"tradeSource"`
	TradeType       string                        `json:"tradeType"`
	FeeCurrency     string                        `json:"feeCurrency"`
	TakerFeeRate    string                        `json:"takerFeeRate"`
	MakerFeeRate    string                        `json:"makerFeeRate"`
	CreatedAt       int64                         `json:"createdAt"`
	Stop            kucoinentity.StopOrderTrigger `json:"stop"`
	StopTriggerTime string                        `json:"stopTriggerTime"`
	StopPrice       string                        `json:"stopPrice"`
}

func (j *stopOrderDetailJSON) toBaseEntity() (*kucoinentity.StopOrderDetailInfo, error) {
	price, err := strconv.ParseFloat(j.Price, 10)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseFloat(j.Size, 10)
	if err != nil {
		return nil, err
	}

	TakerFeeRate, err := strconv.ParseFloat(j.TakerFeeRate, 10)
	if err != nil {
		return nil, err
	}

	MakerFeeRate, err := strconv.ParseFloat(j.MakerFeeRate, 10)
	if err != nil {
		return nil, err
	}

	StopPrice, err := strconv.ParseFloat(j.StopPrice, 10)
	if err != nil {
		return nil, err
	}

	var stopTriggerTime time.Time
	stopTriggerTimeINT, err := strconv.ParseInt(j.StopTriggerTime, 10, 64)
	if err == nil {
		stopTriggerTime = time.Unix(stopTriggerTimeINT, 0)
	}

	entity := kucoinentity.StopOrderDetailInfo{
		Id:              j.Id,
		Symbol:          j.Symbol,
		UserId:          j.UserId,
		Status:          j.Status,
		Type:            j.Type,
		Side:            j.Side,
		Price:           price,
		Size:            size,
		Funds:           j.Funds,
		Stp:             j.Stp,
		TimeInForce:     j.TimeInForce,
		CancelAfter:     time.Unix(j.CancelAfter, 0),
		PostOnly:        j.PostOnly,
		Hidden:          j.Hidden,
		Iceberg:         j.Iceberg,
		VisibleSize:     j.VisibleSize,
		Channel:         j.Channel,
		ClientOid:       j.ClientOid,
		Remark:          j.Remark,
		Tags:            j.Tags,
		OrderTime:       time.Unix(j.OrderTime, 0),
		DomainId:        j.DomainId,
		TradeSource:     j.TradeSource,
		TradeType:       j.TradeType,
		FeeCurrency:     j.FeeCurrency,
		TakerFeeRate:    TakerFeeRate,
		MakerFeeRate:    MakerFeeRate,
		CreatedAt:       time.Unix(j.CreatedAt, 0),
		Stop:            j.Stop,
		StopTriggerTime: stopTriggerTime,
		StopPrice:       StopPrice,
	}

	return &entity, nil
}

func (hc *HTTPClient) StopOrderDetail(orderID string) (*kucoinentity.StopOrderDetailInfo, error) {
	url := strings.Join([]string{stopOrderEndpoint, "/", orderID}, "")

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
		return nil, hc.logAndReturnWrappedErr("stop order detail request err", err)
	}

	b, err := hc.handleResponse(response)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("handle stop order detail response err", err)
	}

	var j = &stopOrderDetailJSON{}

	if err = json.Unmarshal(b, j); err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal stop order detail err", err)
	}

	base, err := j.toBaseEntity()
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("stop order detail parsing err", err)
	}

	return base, nil
}
