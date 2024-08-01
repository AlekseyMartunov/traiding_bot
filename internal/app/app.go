// Package kucoinbots is used to run all kucoin-client bots.
package kucoinbots

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"tradingbot/internal/config"
	kucoinentity "tradingbot/internal/entity"
	httpclient "tradingbot/internal/http/kucoin-client"
	kucoinheader "tradingbot/internal/secrets-header"
	wsClient "tradingbot/internal/websocket/kucoin-client"
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

	webSocketClient, err := wsClient.New(log, conf, kucoinheader.CreateSecretsHeaders, true)
	if err != nil {
		return err
	}
	messages, errors, err := webSocketClient.Connect(ctx)

	err = webSocketClient.Subscribe(
		wsClient.NewAccountBalanceChangeMessageSubscribe(),
		wsClient.NewOrderChangeMessageSubscribe(),
	)
	if err != nil {
		return err
	}

	clientHTTP := httpclient.New(log, &conf.Kucoin, kucoinheader.CreateSecretsHeaders)

	go func() {
		time.Sleep(time.Second * 5)
		order := kucoinentity.MarketOrder{
			OrderID:       "",
			ClientOrderID: strconv.FormatInt(time.Now().UnixNano(), 10),
			Side:          kucoinentity.Buy,
			Funds:         10,
			Pair:          "SOL-USDT",
			Remark:        "",
			Time:          time.Now(),
		}

		err := clientHTTP.PlaceMarketOrder(&order)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("ORDER_ID", order.OrderID)
	}()

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

// export CONFIG_PATH_KUCOIN=config/config.yaml
