package model

import "banks-api/internal/pkg/database/schema"

type Translation struct {
	ID               int64
	Lexeme           string
	TranslatedLexeme string
	SourceLanguage   string
	TargetLanguage   string
}

func ConvertTranslationToModel(tr *schema.Translation) *Translation {
	return &Translation{
		ID:               tr.ID,
		Lexeme:           tr.Lexeme,
		TranslatedLexeme: tr.TranslatedLexeme,
		SourceLanguage:   tr.SourceLanguage,
		TargetLanguage:   tr.TargetLanguage,
	}
}

func ConvertTranslationsToModel(translations []schema.Translation) []*Translation {
	modelTranslations := make([]*Translation, 0, len(translations))

	for _, tr := range translations {
		modelTranslations = append(modelTranslations, ConvertTranslationToModel(&tr))
	}

	return modelTranslations
}
