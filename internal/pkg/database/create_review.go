package database

import (
	"context"
	"time"

	"banks-api/internal/pkg/model"
	"github.com/pkg/errors"
)

func (r *BanksRepository) CreateReview(ctx context.Context, review *model.Review) (int64, error) {
	query := psql.
		Insert(reviewsTableName).
		Columns(
			reviewContentColumn,
			reviewIsApprovedColumn,
			reviewUserEmailColumn,
			reviewUserPhoneColumn,
			reviewRatingColumn,
			reviewBankIDColumn,
			reviewUserNameColumn,
			reviewDateColumn,
		).
		Values(
			review.Content,
			review.IsApproved,
			review.UserEmail,
			review.UserPhone,
			review.Rating,
			review.BankID,
			review.UserName,
			time.Now(),
		).
		Suffix(`RETURNING id;`)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "CreateReview.ToSql")
	}

	var reviewID int64
	if err = r.pool.QueryRow(ctx, rawSQL, args...).Scan(&reviewID); err != nil {
		return 0, errors.Wrap(err, "UpdateBank.Exec")
	}

	return reviewID, nil
}
