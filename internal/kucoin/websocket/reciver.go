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

var notMessageError = errors.New("not a ticker message error")

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

func (r *Receiver) Run(ctx context.Context) <-chan *entity.Ticker {
	ch := make(chan *entity.Ticker)
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
					if !errors.Is(err, notMessageError) {
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

func (r *Receiver) handleMessage(b []byte) (*entity.Ticker, error) {
	t := ticker{}
	json.Unmarshal(b, &t)

	switch t.Type {
	case "pong":
		return nil, notMessageError

	case "ack":
		return nil, notMessageError

	case "message":
		return t.toBaseTicker(), nil
	}

	return nil, notMessageError
}

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
	fmt.Println(configResponse)
	r.log.Info("kucoin websocket config successfully received")
	return nil
}

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

	r.log.Info("websocket connection for kucoin successfully created")
	return nil
}

func (r *Receiver) subscribe() error {
	sub := subscribeMessage{
		Id:             1234567890,
		Type:           "subscribe",
		Topic:          fmt.Sprintf("%s:%s", "/order/ticker", strings.Join(r.pairs, ",")),
		PrivateChannel: false,
		Response:       false,
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
	r.log.Info("kucoin websocket successfully subscribed")

	return nil
}
