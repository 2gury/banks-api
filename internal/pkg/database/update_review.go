package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r *BanksRepository) UpdateReview(ctx context.Context, reviewID int64, isApproved bool) error {
	query := psql.
		Update(reviewsTableName).
		Set(reviewIsApprovedColumn, isApproved).
		Where(sq.Eq{reviewIDColumn: reviewID})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateReview.ToSql")
	}

	if _, err = r.pool.Exec(ctx, rawSQL, args...); err != nil {
		return errors.Wrap(err, "UpdateReview.Exec")
	}

	return nil
}
