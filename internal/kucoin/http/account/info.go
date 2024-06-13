package kucoinaccount

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"
	kucoinheader "tradingbot/internal/kucoin/http/header"
)

// GetAccountInfo returns info about all currency accounts.
func (am *AccountManager) GetAccountInfo() ([]*kucoinentity.AccountInfo, error) {
	headers := kucoinheader.CreateSecretsHeaders(
		http.MethodGet,
		accountEndpoint,
		"",
		am.cfg.Secret(),
		am.cfg.PassPhrase(),
		am.cfg.Key(),
		am.cfg.Version(),
		time.Now(),
	)

	resp, err := am.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{am.cfg.BaseEndpoint(), accountEndpoint}, ""))

	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", resp.String(), resp.StatusCode()))
		return nil, kucoinerrors.ErrStatusCodeIsNot200
	}

	b := bytes.TrimPrefix(resp.Body(), []byte(`{"code":"200000","data":`))
	b = bytes.TrimSuffix(b, []byte("}"))

	jsonArr := make([]*accountInfoJSON, 0, 0)
	err = json.Unmarshal(b, &jsonArr)
	if err != nil {
		am.log.Error(err.Error())
		return nil, kucoinerrors.ErrUnmarshal
	}

	resultArr := make([]*kucoinentity.AccountInfo, 0, len(jsonArr))

	for _, v := range jsonArr {
		e, err := v.toBaseEntity()
		if err != nil {
			am.log.Error(err.Error())
			return nil, kucoinerrors.ErrRecastDTO
		}
		resultArr = append(resultArr, e)
	}

	return resultArr, nil
}
