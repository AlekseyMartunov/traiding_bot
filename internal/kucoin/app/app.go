// Package kucoinbots is used to run all kucoin bots.
package kucoinbots

import (
	"context"
	"fmt"
	"tradingbot/internal/kucoin/config"
	"tradingbot/internal/kucoin/websocket/exchange"
	ml "tradingbot/internal/kucoin/websocket/prediction"
	"tradingbot/pkg/logger"
)

func Run(ctx context.Context) error {
	//logger, err := l.NewLogger("trace", "127.0.0.1:5170", true)
	log, err := logger.New(nil)
	if err != nil {
		return fmt.Errorf("creation logger error: %w", err)
	}

	conf := config.NewConfig()
	err = conf.ParseEnvironment()
	if err != nil {
		return fmt.Errorf("parse config error: %w", err)
	}

	//accountManager := kucoinaccount.New(logger, conf)
	//accountManager.GetAccountInfo()

	//pool, err := pgxpool.New(ctx, "postgres://test:test@localhost:5432/test")
	//if err != nil {
	//	return fmt.Errorf("connect to db error: %w", err)
	//}

	//storage := postgresRepo.NewStorage(pool)
	//test := kucoinEntity.OrderDetailInfo{
	//	Id:          "48jdshdjsnd",
	//	Side:        "Sell",
	//	ClientOid:   "clientoid",
	//	Symbol:      "eth-BTC",
	//	Funds:       "200",
	//	Fee:         "123",
	//	FeeCurrency: "USTD",
	//	CreatedAt:   time.Now(),
	//}
	//
	//err = storage.CloseMarketPosition(ctx, "bot1", &test)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//orderManager := kucoinorders.NewKucoinOrderManager(logger, conf)
	//i, err := orderManager.GetOrderDetail("")
	//if err != nil {
	//	fmt.Println(i)
	//	logger.Info(err.Error())
	//}
	//fmt.Println(i)

	//accountManager := kucoinaccount.New(logger, conf)

	//order := kucoinentity.MarketOrder{
	//	OrderID:       "",
	//	ClientOrderID: "84584",
	//	Side:          kucoinentity.Sell,
	//	Funds:         6.45,
	//	Size:          0,
	//	Pair:          "WEST-USDT",
	//	Time:          time.Time{},
	//}
	//
	//orderManager := kucoinorders.NewKucoinOrderManager(logger, conf)
	//
	//err = orderManager.PlaceMarketOrder(&order)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(orderManager.GetCurrencyConfig("WEST-USDT"))

	mlService, err := ml.New(log, conf)
	if err != nil {
		return err
	}

	kucoinWSReceiver, err := kucoinreceiver.NewReceiver("", log, []string{"BTC-USDT", "ETH-USDT", "SOL-USDT"})
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("creation kucoin websocket reciver error: %w", err)
	}
	ch := kucoinWSReceiver.Run(ctx)

	go func() {
		for {
			res, err := mlService.ReadMessage()
			fmt.Println("result:", res)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)
		}
	}()

	for {
		select {
		case <-ctx.Done():
		case t := <-ch:
			err = mlService.SendMessage(t)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}
