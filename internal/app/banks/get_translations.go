package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) GetTranslations(ctx context.Context, _ *banks.GetTranslationsRequest) (*banks.GetTranslationsResponse, error) {
	translations, err := w.service.GetTranslations(ctx)
	if err != nil {
		return nil, err
	}

	return &banks.GetTranslationsResponse{
		Translations: mappers.ModelsToPBTranslations(translations),
	}, nil
}
