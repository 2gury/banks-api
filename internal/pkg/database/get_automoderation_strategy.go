package database

import (
	"banks-api/internal/pkg/database/schema"
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetAutomoderationStrategy(ctx context.Context) (*schema.Settings, error) {
	query := psql.
		Select(
			settingsIDColumn,
			settingsAutomoderationStrategyColumn,
		).From(settingsTableName)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBanks.ToSql")
	}

	var res []schema.Settings
	err = pgxscan.Select(ctx, r.pool, &res, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "GetTranslationByText.Get")
	}

	if len(res) > 0 {
		return &res[0], nil
	}

	return nil, nil
}
