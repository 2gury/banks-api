package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) GetPossibleBanks(ctx context.Context, _ *banks.GetPossibleBanksRequest) (*banks.GetPossibleBanksResponse, error) {
	possibleBanks, err := w.service.GetPossibleBanks(ctx)
	if err != nil {
		return nil, err
	}

	return mappers.ModelToPBPossibleBanks(possibleBanks), nil
}
