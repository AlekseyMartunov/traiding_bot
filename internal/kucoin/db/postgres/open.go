package postgresrepo

import (
	"context"

	kucoinentity "tradingbot/internal/kucoin/entity"

	"github.com/jackc/pgx/v5"
)

// OpenMarketPosition save first opening order to db.
func (s *Storage) OpenMarketPosition(ctx context.Context, botName string, o *kucoinentity.OrderDetailInfo) error {
	query1 := `INSERT INTO kucoin_market_order 
    		(order_id, client_order_id, symbol, side, funds, commission, commission_currency, created_time)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id;`

	query2 := `INSERT INTO kucoin_position (symbol, bot_name, fk_open_order_id) VALUES ($1, $2, $3);`

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var id int64
	row := tx.QueryRow(ctx, query1, o.Id, o.ClientOid, o.Symbol, o.Side, o.Funds, o.Fee, o.FeeCurrency, o.CreatedAt)
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query2, o.Symbol, botName, id)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}
