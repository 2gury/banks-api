package banks

import (
	"context"

	"banks-api/internal/pkg/model"
)

func (w *BanksHandler) GetBanks(ctx context.Context, filters *model.BankFilters) ([]*model.Bank, error) {
	banks, err := w.repo.GetBanks(ctx, filters)
	if err != nil {
		return nil, err
	}

	return model.ConvertBanksToModel(banks)
}
