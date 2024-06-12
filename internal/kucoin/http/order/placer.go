package kucoinorders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tradingbot/internal/kucoin/errors"

	"tradingbot/internal/kucoin/entity"
)

func (om *KucoinOrderManager) PlaceMarketOrder(order *kucoinentity.MarketOrder) error {
	body := marketOrderJSON{
		ClientOrderID: order.ClientOrderID,
		Side:          order.Side,
		Symbol:        order.Pair,
		OrderType:     "market",
		Funds:         strconv.FormatFloat(order.Funds, 'f', 6, 64),
	}

	b, err := json.Marshal(body)
	if err != nil {
		om.log.Error(err.Error())
		return err
	}

	headers := om.createHeaders(http.MethodPost, endpoint, string(b))

	response, err := om.client.R().
		SetBody(b).
		SetHeaders(headers).
		Post(strings.Join([]string{baseEndpoint, endpoint}, ""))

	if err != nil {
		om.log.Error(err.Error())
		return err
	}

	if response.StatusCode() != 200 {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return kucoinerrors.StatusCodeIsNot200
	}

	respOrder := responseOrderJSON{}

	err = json.Unmarshal(response.Body(), &respOrder)
	if err != nil {
		om.log.Error(err.Error())
		return err
	}

	if respOrder.Code != successfulCode {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return kucoinerrors.StatusCodeIsNot200
	}

	order.OrderID = respOrder.Data.OrderId

	return nil
}
