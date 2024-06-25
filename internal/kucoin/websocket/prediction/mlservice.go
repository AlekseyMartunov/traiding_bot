package ml

import (
	"encoding/json"
	kucoinentity "tradingbot/internal/kucoin/entity"

	"github.com/gorilla/websocket"
)

type config interface {
	MlServiceAddr() string
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Event(msg string, args ...any)
	Error(msg string, args ...any)
}

type PredictionService struct {
	conn *websocket.Conn
	log  logger
	cfg  config
}

func New(l logger, c config) (*PredictionService, error) {
	wsConn, _, err := websocket.DefaultDialer.Dial(c.MlServiceAddr(), nil)
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

func (ps *PredictionService) SendMessage(ticker *kucoinentity.Ticker) error {
	dto := tickerDTO{}
	dto.fromEntity(ticker)

	b, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	err = ps.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PredictionService) ReadMessage() (*kucoinentity.MlResult, error) {
	_, b, err := ps.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	var dto mlResultDTO

	err = json.Unmarshal(b, &dto)
	if err != nil {
		return nil, err
	}

	return dto.toEntity(), nil
}
