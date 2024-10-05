package database

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type BanksRepository struct {
	pool *pgxpool.Pool
}

func NewBanksRepository(pool *pgxpool.Pool) *BanksRepository {
	return &BanksRepository{
		pool: pool,
	}
}
