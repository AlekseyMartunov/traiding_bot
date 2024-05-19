package infra

import (
	"context"
	"fmt"

	"tradingbot/internal/kucoin/websocket"
	"tradingbot/pkg/tcplogger"
)

func RunApp(ctx context.Context) error {
	logger, err := tcplogger.NewLogger("trace", "127.0.0.1:5170", true)
	if err != nil {
		return fmt.Errorf("creation logger error: %w", err)
	}
	defer logger.Close()

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
