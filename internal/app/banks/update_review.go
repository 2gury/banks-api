package banks

import (
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) UpdateReview(ctx context.Context, req *banks.UpdateReviewRequest) (*banks.UpdateReviewResponse, error) {
	if err := w.service.UpdateReview(ctx, req.GetId(), req.GetIsApproved()); err != nil {
		return nil, err
	}
	
	return &banks.UpdateReviewResponse{}, nil
}
