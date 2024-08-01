package wsclient

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	entity "tradingbot/internal/entity"
)

func NewTickerSubscribeMessages(pairs ...string) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", ticker, strings.Join(pairs, ",")),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewTickerUnSubscribeMessages(pairs ...string) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", ticker, strings.Join(pairs, ",")),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewTickerAllSubscribeMessages() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", ticker, "all"),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewTickerAllUnSubscribeMessages() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", ticker, "all"),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewCandleSubscribeMessages(pair string, candle entity.CandlePeriod) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", candles, strings.Join([]string{pair, string(candle)}, "_")),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewCandleUnSubscribeMessages(pair string, candle entity.CandlePeriod) subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          fmt.Sprintf("%s:%s", candles, strings.Join([]string{pair, string(candle)}, "_")),
		PrivateChannel: false,
		Response:       true,
	}
}

func NewOrderChangeMessageSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          privetOrderChange,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewOrderChangeMessageUnSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          privetOrderChange,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewOrderChangeV2MessageSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          privetOrderChangeV2,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewOrderChangeV2MessageUnSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          privetOrderChangeV2,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewAccountBalanceChangeMessageSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          accountBalanceChange,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewAccountBalanceChangeMessageUnSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          accountBalanceChange,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewStopOrderEventMessageSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           SubscribeMessage,
		Topic:          stopOrderEvent,
		PrivateChannel: true,
		Response:       true,
	}
}

func NewStopOrderEventMessageUnSubscribe() subscribeMessage {
	return subscribeMessage{
		Id:             strconv.FormatInt(time.Now().UnixNano(), 10),
		Type:           UnSubscribeMessage,
		Topic:          stopOrderEvent,
		PrivateChannel: true,
		Response:       true,
	}
}
