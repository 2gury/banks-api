package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) CreateTranslationText(ctx context.Context, req *banks.CreateTranslationTextRequest) (*banks.CreateTranslationTextResponse, error) {
	tr, err := w.service.CreateTranslation(ctx, mappers.PBToModelTranslation(req.GetTranslation()))
	if err != nil {
		return nil, err
	}

	return &banks.CreateTranslationTextResponse{
		Translation: mappers.ModelToPBTranslation(tr),
	}, nil
}
