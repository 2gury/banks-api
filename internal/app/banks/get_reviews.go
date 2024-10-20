package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) GetReviews(ctx context.Context, _ *banks.GetReviewsRequest) (*banks.GetReviewsResponse, error) {
	reviews, err := w.service.GetReviews(ctx)
	if err != nil {
		return nil, err
	}

	return &banks.GetReviewsResponse{
		Reviews: mappers.ModelsToPBReviews(reviews),
	}, nil
}
