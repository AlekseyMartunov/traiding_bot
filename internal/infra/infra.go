// Package infra is used to combine other application packages
//
// Right now I have packages to work only with kucoin.
package infra

import (
	"context"

	kucoinBots "tradingbot/internal/kucoin/app"
)

func RunApp(ctx context.Context) error {
	return kucoinBots.Run(ctx)
}
