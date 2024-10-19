package model

import (
	"encoding/json"

	"banks-api/internal/pkg/database/schema"
)

type Review struct {
	ID        int64
	ProductID int64
	Content   string
	Rating    int

	ReviewData ReviewData

	Approved bool
}

type ReviewData struct {
	PeriodFrom int64   `json:"period_from"`
	PeriodTo   int64   `json:"period_to"`
	RateFrom   float64 `json:"rate_from"`
	RateTo     float64 `json:"rate_to"`
}

func ConvertReviewsToModel(reviews []schema.Review) ([]*Review, error) {
	modelReviews := make([]*Review, 0, len(reviews))

	for _, r := range reviews {
		modelReview, err := ConvertReviewToModel(r)
		if err != nil {
			return nil, err
		}

		modelReviews = append(modelReviews, modelReview)
	}

	return modelReviews, nil
}

func ConvertReviewToModel(review schema.Review) (*Review, error) {
	var reviewData ReviewData
	if err := json.Unmarshal([]byte(review.ReviewData), &reviewData); err != nil {
		return nil, err
	}

	return &Review{
		ID:         review.ID,
		ProductID:  review.ProductID,
		Content:    review.Content,
		Rating:     review.Rating,
		ReviewData: ReviewData(reviewData),
	}, nil
}
