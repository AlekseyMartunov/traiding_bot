package kucoinaccount

import (
	"encoding/json"
	"fmt"
	"strings"

	kucoinentity "tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"
)

// GetCurrencyConfig allows to you to get info on purchase and sales volumes.
func (am *AccountManager) GetCurrencyConfig(currencyPair string) (*kucoinentity.CurrencyConfig, error) {
	url := strings.Join([]string{baseEndpoint, symbolListEndpoint, "/", currencyPair}, "")
	response, err := am.client.R().
		Get(url)

	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	var info currencyConfigJSON

	err = json.Unmarshal(response.Body(), &info)
	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if info.Code != successfulCode {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", response.String(), response.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	return info.toBaseEntity(), nil
}
