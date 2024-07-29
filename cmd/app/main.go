package main

import (
	"context"
	"os"
	"os/signal"
	"tradingbot/internal/infra"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := infra.RunApp(ctx)
	if err != nil {
		panic(err)
	}
}
