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

func (om *KucoinOrderManager) PlaceMarketOrder(order *entity.MarketOrder) error {
	body := marketOrderForRequest{
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
	fmt.Println(string(b))

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
		om.log.Info(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return kucoinerrors.StatusCodeIsNot200
	}

	respOrder := responseOrder{}

	err = json.Unmarshal(response.Body(), &respOrder)
	if err != nil {
		om.log.Error(err.Error())
		return err
	}

	order.OrderID = respOrder.Data.OrderId
	fmt.Println(response.String())

	return nil
}
