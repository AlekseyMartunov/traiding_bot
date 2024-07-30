package httpclient

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tradingbot/internal/entity"
)

type accountInfoJSON struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

func (aj *accountInfoJSON) toBaseEntity() (*kucoinentity.AccountInfo, error) {
	balance, err := strconv.ParseFloat(aj.Balance, 10)
	if err != nil {
		return nil, err
	}

	available, err := strconv.ParseFloat(aj.Available, 10)
	if err != nil {
		return nil, err
	}

	holds, err := strconv.ParseFloat(aj.Holds, 10)
	if err != nil {
		return nil, err
	}

	return &kucoinentity.AccountInfo{
		ID:        aj.Id,
		Currency:  aj.Currency,
		Type:      aj.Type,
		Balance:   balance,
		Available: available,
		Holds:     holds,
	}, nil
}

func (hc *HTTPClient) AccountInfo() ([]*kucoinentity.AccountInfo, error) {
	headers := hc.secretHeaders(
		http.MethodGet,
		accountEndpoint,
		"",
		hc.cfg.Secret(),
		hc.cfg.PassPhrase(),
		hc.cfg.Key(),
		hc.cfg.Version(),
		time.Now(),
	)

	resp, err := hc.client.R().
		SetHeaders(headers).
		Get(strings.Join([]string{hc.cfg.BaseEndpoint(), accountEndpoint}, ""))

	if err != nil {
		return nil, hc.logAndReturnWrappedErr("create response account info err", err)
	}

	b, err := hc.handleResponse(resp)
	if err != nil {
		return nil, hc.logAndReturnWrappedErr("handle response account info err", err)
	}

	jsonArr := make([]*accountInfoJSON, 0, 10)

	if err = json.Unmarshal(b, &jsonArr); err != nil {
		return nil, hc.logAndReturnWrappedErr("unmarshal response body err", err)
	}

	entityArr := make([]*kucoinentity.AccountInfo, 0, len(jsonArr))
	for _, j := range jsonArr {
		e, err := j.toBaseEntity()
		if err != nil {
			return nil, hc.logAndReturnWrappedErr("account info parsing err", err)
		}
		entityArr = append(entityArr, e)
	}

	return entityArr, nil
}
