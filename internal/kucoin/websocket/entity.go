package kucoinreceiver

import (
	"strings"
	"tradingbot/internal/basicentity/entity"
)

// WsConfigResponse type is returned by the exchange
// and contains the configuration info for connecting via websocket
type wsConfigResponse struct {
	Code string `json:"code"`
	Data struct {
		Token           string `json:"token"`
		InstanceServers []struct {
			Endpoint     string `json:"endpoint"`
			Encrypt      bool   `json:"encrypt"`
			Protocol     string `json:"protocol"`
			PingInterval int    `json:"pingInterval"`
			PingTimeout  int    `json:"pingTimeout"`
		} `json:"instanceServers"`
	} `json:"data"`
}

// To prevent the TCP link being disconnected by the server, the client side needs to send ping messages
// every pingInterval time to the server to keep alive the link. After the ping message is sent
// to the server, the system would return a pong message to the client side. If the server has not
// received any message from the client for a long time, the connection will be disconnected.
// To subscribe channel messages from a certain server, the client side should send subscription message to the server.
//
// If the subscription succeeds, the system will send ack messages to you,
// when the response is set as true.
type infoWSMessage struct {
	Id string `json:"id"`

	// "ping" or "pong" or "ack"
	Type string `json:"type"`
}

type subscribeMessage struct {
	// The id should be an unique value
	Id int64 `json:"id"`

	// "unsubscribe" or "subscribe"
	Type string `json:"type"`

	// Topic needs to be unsubscribed or subscribe. Some topics support to divisional
	// unsubscribe the information of multiple trading pairs through ",".
	Topic string `json:"topic"`

	PrivateChannel bool `json:"privateChannel"`

	// Whether the server needs to return the receipt information of this subscription or not.
	// Set as false by default.
	Response bool `json:"response"`
}

// ticker contains info about pair
type ticker struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  struct {
		BestAsk     string `json:"bestAsk"`
		BestAskSize string `json:"bestAskSize"`
		BestBid     string `json:"bestBid"`
		BestBidSize string `json:"bestBidSize"`
		Price       string `json:"price"`
		Sequence    string `json:"sequence"`
		Size        string `json:"size"`
		Time        int64  `json:"time"`
	} `json:"data"`
	Subject string `json:"subject"`
}

func (t *ticker) toBaseTicker() *entity.Ticker {
	p := strings.Split(t.Topic, ":")
	pair := ""

	if len(p) > 1 {
		pair = p[1]
	}

	base := entity.Ticker{
		Pair:  pair,
		Price: t.Data.Price,
	}

	return &base
}
