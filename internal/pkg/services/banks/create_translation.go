package banks

import (
	"banks-api/internal/pkg/model"
	"context"
)

func (w *BanksHandler) CreateTranslation(ctx context.Context, translation *model.Translation) (*model.Translation, error) {
	translationID, err := w.repo.CreateTranslation(ctx, translation)
	if err != nil {
		return nil, err
	}

	return &model.Translation{
		ID:               translationID,
		Lexeme:           translation.Lexeme,
		TranslatedLexeme: translation.TranslatedLexeme,
		SourceLanguage:   translation.SourceLanguage,
		TargetLanguage:   translation.TargetLanguage,
	}, nil
}
