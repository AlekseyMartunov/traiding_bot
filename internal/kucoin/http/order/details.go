package kucoinorders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	kucoinentity "tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"
	"tradingbot/internal/kucoin/header"
)

// GetOrderDetail allows you to get info about the completed order.
func (om *KucoinOrderManager) GetOrderDetail(orderID string) (*kucoinentity.OrderDetailInfo, error) {
	url := strings.Join([]string{endpoint, "/", orderID}, "")

	headers := kucoinheader.CreateSecretsHeaders(
		http.MethodGet,
		url,
		"",
		om.cfg.Secret(),
		om.cfg.PassPhrase(),
		om.cfg.Key(),
		om.cfg.Version(),
		time.Now(),
	)

	response, err := om.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{om.cfg.BaseEndpoint(), url}, ""))

	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.ErrStatusCodeIsNot200
	}

	var info orderDetailInfoJSON

	err = json.Unmarshal(response.Body(), &info)
	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if info.Code != successfulCode {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.ErrStatusCodeIsNot200
	}

	return info.toBaseEntity()
}
