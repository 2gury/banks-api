package mappers

import (
	"banks-api/internal/pkg/model"
	pbbanks "banks-api/pkg/banks"
)

func ModelToPBTranslation(translation *model.Translation) *pbbanks.Translation {
	if translation == nil {
		return nil
	}

	return &pbbanks.Translation{
		Id:               translation.ID,
		Lexeme:           translation.Lexeme,
		TranslatedLexeme: translation.TranslatedLexeme,
		SourceLanguage:   translation.SourceLanguage,
		TargetLanguage:   translation.TargetLanguage,
	}
}

func ModelsToPBTranslations(translations []*model.Translation) []*pbbanks.Translation {
	if translations == nil {
		return nil
	}

	pbTranslations := make([]*pbbanks.Translation, 0, len(translations))

	for _, tr := range translations {
		pbTranslations = append(pbTranslations, ModelToPBTranslation(tr))
	}

	return pbTranslations
}

func PBToModelTranslation(translation *pbbanks.Translation) *model.Translation {
	if translation == nil {
		return nil
	}

	return &model.Translation{
		ID:               translation.Id,
		Lexeme:           translation.Lexeme,
		TranslatedLexeme: translation.TranslatedLexeme,
		SourceLanguage:   translation.SourceLanguage,
		TargetLanguage:   translation.TargetLanguage,
	}
}
