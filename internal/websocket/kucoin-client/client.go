package wsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

const (
	websocketEndpoint = "/api/v1/bullet-public"
)

const (
	defaultTimeout = time.Second * 10
)

type httpConfig struct {
	Code string        `json:"code"`
	Data *tokenMessage `json:"data"`
}

type tokenMessage struct {
	Token           string            `json:"token"`
	InstanceServers []*instanceServer `json:"instanceServers"`
}

func (t *tokenMessage) RandomServer() (*instanceServer, error) {
	l := len(t.InstanceServers)
	switch l {
	case 0:
		return nil, errors.New("no available server")
	case 1:
		return t.InstanceServers[0], nil
	default:
		return t.InstanceServers[rand.Intn(l)], nil
	}
}

type instanceServer struct {
	Endpoint     string `json:"endpoint"`
	Encrypt      bool   `json:"encrypt"`
	Protocol     string `json:"protocol"`
	PingInterval int    `json:"pingInterval"`
	PingTimeout  int    `json:"pingTimeout"`
}

type rawMessage struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Sn      string          `json:"sn"`
	Topic   string          `json:"topic"`
	Subject string          `json:"subject"`
	RawData json.RawMessage `json:"data"`
}

const (
	WelcomeMessage     = "welcome"
	PingMessage        = "ping"
	PongMessage        = "pong"
	SubscribeMessage   = "subscribe"
	AckMessage         = "ack"
	UnsubscribeMessage = "unsubscribe"
	ErrorMessage       = "error"
	Message            = "message"
	Notice             = "notice"
	Command            = "command"
)

type wsMessageResponse struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type subscribeMessage struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}

func (s *subscribeMessage) toJsonByte() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return []byte("")
	}
	return b
}

func newPingMessage() *wsMessageResponse {
	return &wsMessageResponse{
		Id:   strconv.FormatInt(time.Now().UnixNano(), 10),
		Type: PingMessage,
	}
}

func NewTickerSubscribeMessages(pairs ...string) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          fmt.Sprintf("/market/ticker:%s", strings.Join(pairs, ",")),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewTickerUnSubscribeMessages(pairs ...string) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnsubscribeMessage,
		Topic:          fmt.Sprintf("/market/ticker:%s", strings.Join(pairs, ",")),
		PrivateChannel: false,
		Response:       true,
	}
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type config interface {
	GetBaseEndpoint() string
}

type WsClient struct {
	wg          *sync.WaitGroup
	errors      chan error
	done        chan struct{}
	pongs       chan string
	ack         chan string
	send        chan string
	rawMessages chan *rawMessage
	conn        *websocket.Conn
	log         logger
	cfg         config
	token       string
	server      *instanceServer
	timeout     time.Duration
}

func New(l logger, c config) (*WsClient, error) {
	ws := WsClient{
		cfg: c,
		log: l,

		errors:      make(chan error),
		done:        make(chan struct{}),
		pongs:       make(chan string),
		ack:         make(chan string),
		send:        make(chan string),
		rawMessages: make(chan *rawMessage, 1024),

		wg:      &sync.WaitGroup{},
		timeout: defaultTimeout,
	}

	err := ws.getConfigMessage()
	if err != nil {
		return nil, err
	}

	return &ws, nil
}

func (ws *WsClient) getConfigMessage() error {
	url := strings.Join([]string{ws.cfg.GetBaseEndpoint(), websocketEndpoint}, "")

	resp, err := resty.New().R().Post(url)
	if err != nil {
		ws.log.Error(err.Error())
		return err
	}

	var hConfig = httpConfig{}

	err = json.Unmarshal(resp.Body(), &hConfig)
	if err != nil {
		ws.log.Error(fmt.Errorf("unmarshal config message error: %w", err).Error())
		return err
	}

	server, err := hConfig.Data.RandomServer()
	if err != nil {
		ws.log.Error(err.Error())
		return err
	}

	ws.server = server
	ws.token = hConfig.Data.Token
	ws.log.Debug("ws client get config message")
	return nil

}

func (ws *WsClient) Connect(ctx context.Context) (<-chan *rawMessage, <-chan error, error) {
	err := ws.createConnection()
	if err != nil {
		ws.log.Error(fmt.Errorf("error creating connection: %w", err).Error())
		return ws.rawMessages, ws.errors, err
	}

	// must read first message
	for {
		m := &rawMessage{}
		if err := ws.conn.ReadJSON(m); err != nil {
			ws.log.Error(err.Error())
			return ws.rawMessages, ws.errors, err
		}

		ws.log.Debug(fmt.Sprintf(
			"Recive a first message ID: %s, type: %s",
			m.ID, m.Type),
		)

		if m.Type == WelcomeMessage {
			break
		}

		if m.Type == ErrorMessage {
			return ws.rawMessages, ws.errors, errors.New(
				fmt.Sprintf("recive a fisrst message: %s", toStringFromJson(m)))
		}

		return ws.rawMessages, ws.errors, errors.New(
			fmt.Sprintf("unknowen type of message: %s", toStringFromJson(m)))
	}

	ws.wg.Add(3)

	go func() {
		ws.wg.Wait()
		close(ws.rawMessages)
		close(ws.errors)
		ws.log.Info("ws client is closing")
	}()

	go ws.write(ctx)
	go ws.handleMessage(ctx)
	go ws.pingMessage(ctx)

	return ws.rawMessages, ws.errors, nil
}

func (ws *WsClient) createConnection() error {
	endpointWithToken := strings.Join([]string{ws.server.Endpoint, "?token=", ws.token}, "")
	conn, _, err := websocket.DefaultDialer.Dial(endpointWithToken, nil)
	if err != nil {
		return err
	}

	ws.conn = conn
	return nil
}

func (ws *WsClient) handleMessage(ctx context.Context) {
	defer func() {
		ws.wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return

		default:
			m := &rawMessage{}
			if err := ws.conn.ReadJSON(m); err != nil {
				ws.log.Error(err.Error())
				ws.errors <- err
			}

			switch m.Type {
			case WelcomeMessage:
			case PongMessage:
				ws.pongs <- m.ID

			case ErrorMessage:
				ws.errors <- errors.New(fmt.Sprintf("Error message: %s", toStringFromJson(m)))

			case AckMessage:
				ws.ack <- m.ID

			case Message, Notice, Command:
				ws.rawMessages <- m

			default:
				ws.errors <- errors.New(fmt.Sprintf("unkown type of message: %s", toStringFromJson(m)))
			}
		}
	}

}

func (ws *WsClient) pingMessage(ctx context.Context) {
	t := time.NewTicker(time.Duration(ws.server.PingInterval/2) * time.Millisecond)
	defer func() {
		ws.wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case <-t.C:
			p := newPingMessage()
			m := toStringFromJson(p)
			ws.log.Debug(fmt.Sprintf("send a ping message: %s", m))

			ws.send <- m

			select {
			case pid := <-ws.pongs:
				if pid != p.Id {
					ws.errors <- errors.New(
						fmt.Sprintf("invalid pong id expected: %s, actual: %s",
							p.Id, pid))
				}

			case <-time.After(time.Duration(ws.server.PingTimeout) * time.Millisecond):
				ws.errors <- errors.New("timeout pong message")
			}
		}
	}

}

func (ws *WsClient) Subscribe(messages ...subscribeMessage) {
	for _, m := range messages {
		ws.log.Debug(fmt.Sprintf("send message: %s", toStringFromJson(m)))

		ws.send <- toStringFromJson(m)

		select {
		case id := <-ws.ack:
			if id != m.Id {
				ws.errors <- errors.New(fmt.Sprintf("invalid received ack id, expected: %s, actual: %s", m.Id, id))
			}

		case <-time.After(ws.timeout):
			ws.errors <- errors.New("ack message timeout")
		}
	}
}

func (ws *WsClient) write(ctx context.Context) {
	defer func() {
		ws.wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case s := <-ws.send:
			err := ws.conn.WriteMessage(websocket.TextMessage, []byte(s))
			if err != nil {
				ws.errors <- err
			}
		}
	}
}

func toStringFromJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
