package infra

import (
	"context"
	"fmt"
	"tradingbot/pkg/tcplogger"
)

func RunApp(ctx context.Context) error {
	logger, err := tcplogger.NewLogger("trace", "127.0.0.1:5170", true)
	if err != nil {
		return fmt.Errorf("creation logger error: %w", err)
	}
	defer logger.Close()

	return nil
}
