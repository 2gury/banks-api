package model

import (
	"banks-api/internal/pkg/database/schema"
)

type Review struct {
	ID         int64
	Content    string
	IsApproved bool
	UserEmail  string
	UserPhone  string
	Rating     int64
	BankID     int64
}

func ConvertReviewsToModel(reviews []schema.Review) []*Review {
	if reviews == nil {
		return nil
	}

	modelReviews := make([]*Review, 0, len(reviews))

	for _, r := range reviews {
		modelReviews = append(modelReviews, ConvertReviewToModel(r))
	}

	return modelReviews
}

func ConvertReviewToModel(review schema.Review) *Review {
	return &Review{
		ID:         review.ID,
		Content:    review.Content,
		IsApproved: review.IsApproved,
		UserEmail:  review.UserEmail,
		UserPhone:  review.UserPhone,
		Rating:     review.Rating,
		BankID:     review.BankID,
	}
}
