package banks

import (
	"context"

	"banks-api/pkg/banks"
)

func (w *BanksHandler) GetBanks(ctx context.Context, _ *banks.GetBanksRequest) (*banks.GetBanksResponse, error) {
	_, err := w.service.GetBanks(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
