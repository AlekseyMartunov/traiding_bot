package postgresrepo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	kucoinentity "tradingbot/internal/kucoin/entity"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(p *pgxpool.Pool) *Storage {
	return &Storage{pool: p}
}

func (s *Storage) OpenMarketOrder(ctx context.Context, o *kucoinentity.OrderDetailInfo) error {
	query1 := `INSERT INTO kucoin_market_order 
    		(order_id, client_order_id, symbol, side, funds, commission, commission_currency, created_time)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id;`

	query2 := `INSERT INTO kucoin_position (symbol, fk_open_order_id) VALUES ($1, $2);`

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

	fmt.Println(id)

	_, err = tx.Exec(ctx, query2, o.Symbol, id)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}
