package banks

import (
	"context"
)

func (w *BanksHandler) UpdateReview(ctx context.Context, reviewID int64, isApproved bool) error {
	if err := w.repo.UpdateReview(ctx, reviewID, isApproved); err != nil {
		return err
	}

	return nil
}
