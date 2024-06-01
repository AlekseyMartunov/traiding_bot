package infra

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"tradingbot/internal/kucoin/config"
	postgresrepo "tradingbot/internal/kucoin/db/postgres"
	kucoinentity "tradingbot/internal/kucoin/entity"
	"tradingbot/internal/kucoin/websocket"
	"tradingbot/pkg/tcplogger"
)

func RunApp(ctx context.Context) error {
	logger, err := tcplogger.NewLogger("trace", "127.0.0.1:5170", true)
	if err != nil {
		return fmt.Errorf("creation logger error: %w", err)
	}
	defer logger.Close()

	conf := config.NewConfig()
	err = conf.ParseEnvironment()
	if err != nil {
		return err
	}

	//orderManager := kucoinorders.NewKucoinOrderManager(logger, conf)
	//i, err := orderManager.GetOrderDetail("")
	//if err != nil {
	//	fmt.Println(i)
	//	logger.Info(err.Error())
	//}
	//fmt.Println(i)

	//accountManager := kucoinaccount.NewAccountManager(logger, conf)

	test := kucoinentity.OrderDetailInfo{
		Id:          "48jdshdjsnd",
		Side:        "Sell",
		ClientOid:   "clientoid",
		Symbol:      "eth-BTC",
		Funds:       "200",
		Fee:         "123",
		FeeCurrency: "USTD",
		CreatedAt:   time.Now(),
	}

	dsn := "postgres://test:test@localhost:5432/test"
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	repo := postgresrepo.NewStorage(pool)

	err = repo.OpenMarketOrder(ctx, &test)
	if err != nil {
		fmt.Println(err)
	}

	kucoinWSReceiver, err := kucoinreceiver.NewReceiver("", logger, []string{"BTC-USDT", "ETH-USDT"})
	if err != nil {
		logger.Error(err.Error())
		return fmt.Errorf("creation kucoin websocket reciver error: %w", err)
	}
	ch := kucoinWSReceiver.Run(ctx)

	for {
		select {
		case <-ctx.Done():
		case t := <-ch:
			fmt.Println(t)
		}
	}

	return nil
}
