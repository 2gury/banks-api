package banks

import (
	"banks-api/internal/pkg/model"
	"context"
)

func (w *BanksHandler) GetTranslations(ctx context.Context) ([]*model.Translation, error) {
	schemaTranslations, err := w.repo.GetTranslations(ctx)
	if err != nil {
		return nil, err
	}

	return model.ConvertTranslationsToModel(schemaTranslations), nil
}
