package kucoinreceiver

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

// ping sends a ping message every few seconds to maintain the connection.
func (r *Receiver) ping(ctx context.Context) {
	d := time.Duration((r.config.Data.InstanceServers[0].PingInterval / 1000) / 2)
	ticker := time.NewTicker(d * time.Second)

	b, err := json.Marshal(infoWSMessage{Id: "123", Type: "ping"})
	if err != nil {
		r.log.Error(err.Error())
	}

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				err = r.conn.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					return
				}
			}
		}
	}()
}
