// Package kucoinbots is used to run all kucoin-client bots.
package kucoinbots

import (
	"context"
	"fmt"
	"tradingbot/internal/config"
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

	log.Info("HI")

	return nil
}

// export CONFIG_PATH_KUCOIN=config/config.yaml
