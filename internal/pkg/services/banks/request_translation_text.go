package banks

import (
	"context"
	"fmt"
)

func (w *BanksHandler) RequestTranslationText(ctx context.Context, language, text string) (string, error) {
	if contains(language, kzSymbols) {
		translatedText, err := w.translatorService.TranslateFromRFToKZText(ctx, text)
		if err != nil {
			return "", err
		}

		return translatedText, nil
	}

	return "", fmt.Errorf("unknown language, you can use `kk`, `Казахский`, `Казахстан`, `КЗ`, `кз`, `Kazakh`, `kazakhstan`")
}
