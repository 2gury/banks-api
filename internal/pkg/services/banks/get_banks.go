package banks

import (
	"context"
	
	"banks-api/internal/pkg/model"
)

func (w *BanksHandler) GetBanks(ctx context.Context) ([]*model.Bank, error) {
	banks, err := w.repo.GetBanks(ctx)
	if err != nil {
		return nil, err
	}

	return model.ConvertBanksToModel(banks), nil
}
