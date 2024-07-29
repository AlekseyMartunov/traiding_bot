package ml

import (
	"encoding/json"
	"strings"

	kucoinentity "tradingbot/internal/kucoin/entity"

	"github.com/gorilla/websocket"
)

const (
	mlEndpoint = "/v1/simple_average"
)

type config interface {
	GetMlAddr() string
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type PredictionService struct {
	conn *websocket.Conn
	log  logger
	cfg  config
}

func New(l logger, c config) (*PredictionService, error) {
	addr := strings.Join([]string{"ws://", c.GetMlAddr(), mlEndpoint}, "")
	wsConn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}
	l.Info("ml service websocket connection successfully created")
	return &PredictionService{
		conn: wsConn,
		cfg:  c,
		log:  l,
	}, nil
}

func (ps *PredictionService) SendTickerMessage(ticker *kucoinentity.Ticker) error {
	j := tickerJSON{}
	j.fromEntity(ticker)

	b, err := json.Marshal(j)
	if err != nil {
		return err
	}

	err = ps.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PredictionService) Run() chan kucoinentity.MlServiceRawMessage {
	ch := make(chan kucoinentity.MlServiceRawMessage)
	go func() {
		defer close(ch)

		for {
			_, b, err := ps.conn.ReadMessage()
			if err != nil {
				ps.log.Error(err.Error())
				continue
			}

			var j kucoinentity.MlServiceRawMessage
			err = json.Unmarshal(b, &j)
			if err != nil {
				ps.log.Error(err.Error())
				continue
			}

			ch <- j
		}
	}()

	return ch
}
