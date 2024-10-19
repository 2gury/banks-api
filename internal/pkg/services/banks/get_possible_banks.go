package banks

import (
	"banks-api/internal/pkg/model"
	"context"
)

func (w *BanksHandler) GetPossibleBanks(_ context.Context) (*model.OffersResponse, error) {
	return w.leadgidService.GetPossibleBanks()
}
