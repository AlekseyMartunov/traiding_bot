package kucoinreceiver

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// createWSConnection create WS connection using token and endpoint from config message.
// After creating the connection server, send a welcome message to test the connection.
func (r *Receiver) createWSConnection() error {
	token := r.config.Data.Token
	endpoint := r.config.Data.InstanceServers[0].Endpoint

	endpointWithToken := strings.Join([]string{endpoint, "?token=", token}, "")
	conn, _, err := websocket.DefaultDialer.Dial(endpointWithToken, nil)

	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	r.conn = conn

	_, b, err := r.conn.ReadMessage()
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	message := infoWSMessage{}
	err = json.Unmarshal(b, &message)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	if message.Type != "welcome" {
		r.log.Warn(fmt.Sprintf("get not expected message type: id: %s, type: %s",
			message.Id,
			message.Type),
		)
	}

	r.log.Info("kucoin websocket connection successfully created")
	return nil
}
