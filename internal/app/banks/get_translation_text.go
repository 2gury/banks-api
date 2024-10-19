package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) GetTranslationText(ctx context.Context, req *banks.GetTranslationTextRequest) (*banks.GetTranslationTextResponse, error) {
	translation, err := w.service.GetTranslationByText(ctx, req.GetLanguage(), req.GetText())
	if err != nil {
		return nil, err
	}

	return &banks.GetTranslationTextResponse{
		Translation: mappers.ModelToPBTranslation(translation),
	}, nil
}
