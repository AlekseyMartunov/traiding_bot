package main

import (
	"context"
	"tradingbot/internal/infra"
)

func main() {
	err := infra.RunApp(context.Background())
	if err != nil {
		panic(err)
	}
}
