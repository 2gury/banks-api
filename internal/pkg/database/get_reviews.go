package database

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"

	"banks-api/internal/pkg/database/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetReviews(ctx context.Context) ([]schema.Review, error) {
	query := psql.
		Select(
			reviewIDColumn,
			reviewContentColumn,
			reviewIsApprovedColumn,
			reviewUserEmailColumn,
			reviewUserPhoneColumn,
			reviewRatingColumn,
			reviewBankIDColumn,
			reviewUserNameColumn,
			reviewDateColumn,
		).From(reviewsTableName).
		Where(sq.Eq{
			reviewIsApprovedColumn: false,
		}).
		OrderBy(fmt.Sprintf("%s DESC", reviewIDColumn))

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBanks.ToSql")
	}

	var res []schema.Review
	err = pgxscan.Select(ctx, r.pool, &res, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "GetReviews.Select")
	}

	return res, nil
}
