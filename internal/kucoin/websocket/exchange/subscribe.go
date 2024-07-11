package kucoinreceiver

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// subscribe tells the server which pair we need to subscribe to.
func (r *Receiver) subscribe() error {
	sub := subscribeMessage{
		Id:             1234567890,
		Type:           "subscribe",
		Topic:          fmt.Sprintf("%s:%s", "/market/ticker", strings.Join(r.pairs, ",")),
		PrivateChannel: false,
		Response:       true,
	}

	b, err := json.Marshal(sub)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	err = r.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	_, b, err = r.conn.ReadMessage()
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	infoMessage := infoWSMessage{}

	err = json.Unmarshal(b, &infoMessage)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	if infoMessage.Id != "1234567890" && infoMessage.Type != "ack" {
		return subscribeErr
	}

	r.log.Info("kucoin websocket successfully subscribed")

	return nil
}
