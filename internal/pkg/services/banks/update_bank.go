package banks

import (
	"context"

	"banks-api/internal/pkg/model"
)

func (w *BanksHandler) UpdateBank(ctx context.Context, bank *model.Bank) error {
	if err := w.repo.UpdateBank(ctx, bank); err != nil {
		return err
	}

	return nil
}
