package main

import (
	"context"
	"os"
	"os/signal"
	"tradingbot/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := kucoinbots.Run(ctx)
	if err != nil {
		panic(err)
	}
}
