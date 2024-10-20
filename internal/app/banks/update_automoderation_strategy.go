package banks

import (
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) UpdateAutomoderationStrategy(ctx context.Context, req *banks.UpdateAutomoderationStrategyRequest) (*banks.UpdateAutomoderationStrategyResponse, error) {
	if err := w.service.UpdateAutomoderationStrategy(ctx, req.GetAutomoderationEnable()); err != nil {
		return nil, err
	}

	return &banks.UpdateAutomoderationStrategyResponse{}, nil
}
