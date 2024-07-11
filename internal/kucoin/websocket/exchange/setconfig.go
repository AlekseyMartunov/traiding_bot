package kucoinreceiver

import (
	"encoding/json"
	"io"
	"net/http"
)

// setConfigForWSConnection make a http request to get config message for creating ws connection.
func (r *Receiver) setConfigForWSConnection(url string) error {
	configResponse := wsConfigResponse{}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	err = json.Unmarshal(b, &configResponse)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	r.config = configResponse
	r.log.Info("kucoin websocket config successfully received")
	return nil
}
