package banks

import (
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) RequestTranslationText(ctx context.Context, req *banks.RequestTranslationTextRequest) (*banks.RequestTranslationTextResponse, error) {
	translatedText, err := w.service.RequestTranslationText(ctx, req.GetLanguage(), req.GetText())
	if err != nil {
		return nil, err
	}

	return &banks.RequestTranslationTextResponse{
		Language: req.GetLanguage(),
		Text:     translatedText,
	}, nil
}
