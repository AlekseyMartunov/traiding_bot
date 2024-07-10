package kucoinaccount

import (
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
		am.cfg.GetSecret(),
		am.cfg.GetPassPhrase(),
		am.cfg.GetKey(),
		am.cfg.GetVersion(),
		time.Now(),
	)

	resp, err := am.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{am.cfg.GetBaseEndpoint(), accountEndpoint}, ""))

	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", resp.String(), resp.StatusCode()))
		return nil, kucoinerrors.ErrStatusCodeIsNot200
	}

	var j accountInfoJSON
	err = json.Unmarshal(resp.Body(), &j)
	if err != nil {
		am.log.Error(err.Error())
		return nil, err
	}

	if j.Code != successfulCode {
		am.log.Error(fmt.Sprintf("body: %s, code: %d", resp.String(), j.Code))
	}

	return j.toBaseEntity()
}
