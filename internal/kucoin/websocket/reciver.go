// Package kucoinreceiver contains the implementation of the client's
// websocket to obtain prices for pairs.
//
// four steps to set up a websocket connection:
// 1 - make a http request to get config message for creating ws connection.
// 2 - using the token and endpoint from config message to make a WS connection.
// 3 - after a certain time, which is obtained from the config,
// send a ping message to the connection so that the server does not close the connection.
// 4 - to subscribe to receive the necessary information, send the required connection pairs.
package kucoinreceiver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"tradingbot/internal/kucoin/entity"

	"github.com/gorilla/websocket"
)

const (
	configURl = "https://api.kucoin.com/api/v1/bullet-public"
)

var (
	notTickerMessageError = errors.New("not a ticker message error")
	subscribeErr          = errors.New("kucoin ws subscribing error")
)

type logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
}

type Receiver struct {
	conn   *websocket.Conn
	log    logger
	config wsConfigResponse
	pairs  []string
}

func NewReceiver(url string, l logger, pairs []string) (*Receiver, error) {
	if url == "" {
		url = configURl
	}

	r := Receiver{
		log:   l,
		pairs: pairs,
	}

	err := r.setConfigForWSConnection(url)
	if err != nil {
		return nil, err
	}

	err = r.createWSConnection()
	if err != nil {
		return nil, err
	}

	err = r.subscribe()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Run send ticker message in to chan.
func (r *Receiver) Run(ctx context.Context) <-chan *kucoinentity.Ticker {
	ch := make(chan *kucoinentity.Ticker)
	r.ping(ctx)

	go func() {
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				return

			default:
				_, b, err := r.conn.ReadMessage()
				if err != nil {
					r.log.Warn(err.Error())
					break
				}

				t, err := r.handleMessage(b)
				if err != nil {
					if !errors.Is(err, notTickerMessageError) {
						r.log.Warn(err.Error())
					}
					break
				}
				ch <- t
			}
		}
	}()
	return ch
}

// ping sends a ping message every few seconds to maintain the connection.
func (r *Receiver) ping(ctx context.Context) {
	d := time.Duration((r.config.Data.InstanceServers[0].PingInterval / 1000) / 2)
	ticker := time.NewTicker(d * time.Second)

	b, err := json.Marshal(infoWSMessage{Id: "123", Type: "ping"})
	if err != nil {
		r.log.Error(err.Error())
	}

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				err = r.conn.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					return
				}
			}
		}
	}()
}

// handleMessage checks the received message and returns only the ticker message,
// all service messages will be skipped.
//
// The server responds to every ping message,
// so every few seconds we receive a response to a ping message,
// this response will be skipped.
func (r *Receiver) handleMessage(b []byte) (*kucoinentity.Ticker, error) {
	t := ticker{}
	json.Unmarshal(b, &t)

	switch t.Type {
	case "pong":
		return nil, notTickerMessageError

	case "ack":
		return nil, notTickerMessageError

	case "message":
		return t.toBaseTicker(), nil
	}

	return nil, notTickerMessageError
}

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
