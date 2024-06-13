package kucoinorders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tradingbot/internal/kucoin/errors"
	kucoinheader "tradingbot/internal/kucoin/http/header"

	"tradingbot/internal/kucoin/entity"
)

// PlaceMarketOrder takes the MarketOrder struct and places it in kucoin.
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

	headers := kucoinheader.CreateSecretsHeaders(
		http.MethodPost,
		endpoint,
		string(b),
		om.cfg.Secret(),
		om.cfg.PassPhrase(),
		om.cfg.Key(),
		om.cfg.Version(),
		time.Now(),
	)

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
		return kucoinerrors.ErrStatusCodeIsNot200
	}

	respOrder := responseOrderJSON{}

	err = json.Unmarshal(response.Body(), &respOrder)
	if err != nil {
		om.log.Error(err.Error())
		return err
	}

	if respOrder.Code != successfulCode {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return kucoinerrors.ErrStatusCodeIsNot200
	}

	order.OrderID = respOrder.Data.OrderId

	return nil
}
