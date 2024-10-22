package model

import (
	"banks-api/internal/pkg/database/schema"
	"time"
)

type Review struct {
	ID         int64
	Content    string
	IsApproved bool
	UserEmail  string
	UserPhone  string
	Rating     int64
	BankID     int64
	UserName   string
	Date       time.Time
	BankName   string
}

func ConvertReviewsToModel(reviews []schema.Review, banks map[int64]string) []*Review {
	if reviews == nil {
		return nil
	}

	modelReviews := make([]*Review, 0, len(reviews))

	for _, r := range reviews {
		bankName := banks[r.BankID]

		modelReviews = append(modelReviews, ConvertReviewToModel(r, bankName))
	}

	return modelReviews
}

func ConvertReviewToModel(review schema.Review, bankName string) *Review {
	return &Review{
		ID:         review.ID,
		Content:    review.Content,
		IsApproved: review.IsApproved,
		UserEmail:  review.UserEmail,
		UserPhone:  review.UserPhone,
		Rating:     review.Rating,
		BankID:     review.BankID,
		UserName:   review.UserName,
		Date:       review.Date,
		BankName:   bankName,
	}
}
