// Package kucoinbots is used to run all kucoin bots.
package kucoinbots

import (
	"context"
	"encoding/json"
	"fmt"

	"tradingbot/internal/kucoin/config"
	"tradingbot/internal/kucoin/websocket/wsclient"
	"tradingbot/pkg/logger"
)

func Run(ctx context.Context) error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("parse config error: %w", err)
	}

	log, err := logger.New(&conf.Logger)
	if err != nil {
		return fmt.Errorf("creation logger error: %w", err)
	}

	c, err := wsclient.New(log, conf)
	if err != nil {
		fmt.Println(err)
	}
	messages, errors, err := c.Connect(ctx)

	c.Subscribe(wsclient.NewTickerSubscribeMessages("ETH-USDT", "SOL-USDT", "BTC-USDT"))

	for {
		select {
		case m := <-messages:
			b, _ := json.Marshal(m)
			fmt.Println("MESSAGE:", string(b))

		case e := <-errors:
			b, _ := json.Marshal(e)
			fmt.Println("ERROR:", string(b))

		case <-ctx.Done():
			log.Info("app is shutdown")
			return nil
		}

	}
	return nil
}
