package database

import (
	"context"

	"banks-api/internal/pkg/database/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetBanks(ctx context.Context) ([]schema.Bank, error) {
	query := psql.
		Select(bankIDColumn, bankNameColumn).
		From(banksTableName)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBanks.ToSql")
	}

	var res []schema.Bank
	err = pgxscan.Select(ctx, r.pool, &res, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "GetBanks.Select")
	}

	return res, nil
}
