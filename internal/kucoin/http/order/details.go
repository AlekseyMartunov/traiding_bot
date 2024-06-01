package kucoinorders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	kucoinentity "tradingbot/internal/kucoin/entity"

	kucoinerrors "tradingbot/internal/kucoin/errors"
)

func (om *KucoinOrderManager) GetOrderDetail(orderID string) (*kucoinentity.OrderDetailInfo, error) {
	url := strings.Join([]string{endpoint, "/", orderID}, "")
	headers := om.createHeaders(http.MethodGet, url, "")
	response, err := om.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{baseEndpoint, url}, ""))

	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	var info orderDetailInfoJSON

	err = json.Unmarshal(response.Body(), &info)
	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	return info.toBaseEntity()
}
