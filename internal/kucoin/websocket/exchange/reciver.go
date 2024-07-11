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
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
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
		return t.toBaseTicker()
	}

	return nil, notTickerMessageError
}
