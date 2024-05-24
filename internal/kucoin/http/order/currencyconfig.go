package kucoinorders

import (
	"encoding/json"
	"fmt"
	"strings"

	kucoinerrors "tradingbot/internal/kucoin/errors"
)

func (om *KucoinOrderManager) GetCurrencyConfig(currencyPair string) (*currencyConfig, error) {
	url := strings.Join([]string{baseEndpoint, symbolListEndpoint, "/", currencyPair}, "")
	response, err := om.client.R().
		Get(url)

	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		om.log.Info(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	var info currencyConfig

	err = json.Unmarshal(response.Body(), &info)
	if err != nil {
		om.log.Error(err.Error())
		return nil, err
	}

	return &info, nil
}
