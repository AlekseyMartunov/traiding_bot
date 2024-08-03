package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type cancelledOrderIdJSON struct {
	OrderIDs []string `json:"cancelledOrderIds"`
}

func (hc *HTTPClient) CanselStopOrder(orderID string) error {
	url := strings.Join([]string{orderEndpoint, "/", orderID}, "")

	headers := hc.secretHeaders(
		http.MethodDelete,
		url,
		"",
		hc.cfg.Secret(),
		hc.cfg.PassPhrase(),
		hc.cfg.Key(),
		hc.cfg.Version(),
		time.Now(),
	)

	response, err := hc.client.R().
		SetHeaders(headers).
		Delete(strings.Join([]string{hc.cfg.BaseEndpoint(), url}, ""))

	if err != nil {
		return hc.logAndReturnWrappedErr("cansel stop order detail request err", err)
	}

	b, err := hc.handleResponse(response)
	if err != nil {
		return hc.logAndReturnWrappedErr("handle cansel stop order detail response err", err)
	}

	j := &cancelledOrderIdJSON{}

	if err = json.Unmarshal(b, j); err != nil {
		return hc.logAndReturnWrappedErr("unmarshal cansel stop order err", err)
	}

	if j.OrderIDs[0] != orderID {
		return hc.logAndReturnWrappedErr("order ids does not match",
			errors.New(fmt.Sprintf("actual: %s, expected: %s", j.OrderIDs[0], orderID)))
	}

	return nil
}
