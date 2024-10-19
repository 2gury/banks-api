package banks

import (
	"banks-api/internal/pkg/model"
	"context"
	"fmt"
)

func (w *BanksHandler) GetTranslationByText(ctx context.Context, language, text string) (*model.Translation, error) {
	if contains(language, kzSymbols) {
		translatedText, err := w.repo.GetTranslationByText(ctx, "kk", text)
		if err != nil {
			return nil, err
		}

		return model.ConvertTranslationToModel(translatedText), nil
	}

	return nil, fmt.Errorf("unknown language, you can use `kk`, `Казахский`, `Казахстан`, `КЗ`, `кз`, `Kazakh`, `kazakhstan`")
}
