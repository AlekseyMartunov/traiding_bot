package postgresrepo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(p *pgxpool.Pool) *Storage {
	return &Storage{pool: p}
}
