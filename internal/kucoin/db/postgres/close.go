package postgresrepo

import (
	"context"
	kucoinerrors "tradingbot/internal/kucoin/errors"

	kucoinentity "tradingbot/internal/kucoin/entity"

	"github.com/jackc/pgx/v5"
)

// CloseMarketPosition save second closing order in to db.
func (s *Storage) CloseMarketPosition(ctx context.Context, botName string, o *kucoinentity.OrderDetailInfo) error {
	query1 := `INSERT INTO kucoin_market_order 
    			(order_id, client_order_id, symbol, side, funds, commission, commission_currency, created_time)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id;`

	query2 := `UPDATE kucoin_position SET fk_close_order_id = $1 
				WHERE bot_name = $2 AND symbol = $3 AND fk_close_order_id IS NULL;`

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

	result, err := tx.Exec(ctx, query2, id, botName, o.Symbol)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	if result.RowsAffected() == 0 {
		return kucoinerrors.ErrNothingToChange
	}

	return nil
}
