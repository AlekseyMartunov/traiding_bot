package kucoinaccount

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tradingbot/internal/kucoin/entity"

	kucoinerrors "tradingbot/internal/kucoin/errors"
)

func (am *AccountManager) GetAccountInfo() ([]*kucoinentity.AccountInfo, error) {
	headers := am.createHeaders(http.MethodGet, accountEndpoint, "")
	resp, err := am.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{baseEndpoint, accountEndpoint}, ""))

	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", resp.String(), resp.StatusCode()))
		return nil, kucoinerrors.StatusCodeIsNot200
	}

	b := bytes.TrimPrefix(resp.Body(), []byte(`{"code":"200000","data":`))
	b = bytes.TrimSuffix(b, []byte("}"))

	jsonArr := make([]*accountInfoJSON, 0, 0)
	json.Unmarshal(b, &jsonArr)

	resultArr := make([]*kucoinentity.AccountInfo, 0, len(jsonArr))

	for _, v := range jsonArr {
		e, err := v.toBaseEntity()
		if err != nil {
			return nil, fmt.Errorf("parsing to base entity err: %w", err)
		}
		resultArr = append(resultArr, e)
	}

	return resultArr, nil
}
