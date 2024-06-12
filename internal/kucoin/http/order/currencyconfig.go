package kucoinorders

import (
	"encoding/json"
	"fmt"
	"strings"

	kucoinentity "tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"
)

func (om *KucoinOrderManager) GetCurrencyConfig(currencyPair string) (*kucoinentity.CurrencyConfig, error) {
	url := strings.Join([]string{baseEndpoint, symbolListEndpoint, "/", currencyPair}, "")
	response, err := om.client.R().
		Get(url)

	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	var info currencyConfigJSON

	err = json.Unmarshal(response.Body(), &info)
	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if info.Code != successfulCode {
		om.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	return info.toBaseEntity(), nil
}
