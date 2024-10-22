package banks

import (
	"banks-api/internal/pkg/model"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) CreateReview(ctx context.Context, req *banks.CreateReviewRequest) (*banks.CreateReviewResponse, error) {
	if err := w.service.CreateReview(ctx, &model.Review{
		Content:   req.GetContent(),
		UserEmail: req.GetUserEmail(),
		UserPhone: req.GetUserPhone(),
		Rating:    req.GetRating(),
		BankID:    req.GetBankId(),
		UserName:  req.GetUserName(),
	}); err != nil {
		return nil, err
	}

	return &banks.CreateReviewResponse{}, nil
}
