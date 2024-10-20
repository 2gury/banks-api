package banks

import (
	"banks-api/internal/pkg/model"
	"context"
)

func (w *BanksHandler) GetReviews(ctx context.Context) ([]*model.Review, error) {
	reviews, err := w.repo.GetReviews(ctx)
	if err != nil {
		return nil, err
	}

	return model.ConvertReviewsToModel(reviews), nil
}
