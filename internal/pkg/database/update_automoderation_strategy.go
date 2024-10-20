package database

import (
	"context"
	"github.com/pkg/errors"
)

func (r *BanksRepository) UpdateAutomoderationStrategy(ctx context.Context, automoderationEnable bool) error {
	query := psql.
		Update(settingsTableName).
		Set(settingsAutomoderationStrategyColumn, automoderationEnable)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateAutomoderationStrategy.ToSql")
	}

	if _, err = r.pool.Exec(ctx, rawSQL, args...); err != nil {
		return errors.Wrap(err, "UpdateAutomoderationStrategy.Exec")
	}

	return nil
}
