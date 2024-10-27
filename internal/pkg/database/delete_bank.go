package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"

	"github.com/pkg/errors"
)

func (r *BanksRepository) DeleteBank(ctx context.Context, id int64) error {
	query := psql.
		Delete(banksTableName).
		Where(sq.Eq{bankIDColumn: id})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "DeleteBank.ToSql")
	}

	_, err = r.pool.Exec(ctx, rawSQL, args...)
	if err != nil {
		return errors.Wrap(err, "DeleteBank.Exec")
	}

	return nil
}
