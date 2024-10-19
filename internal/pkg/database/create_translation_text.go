package database

import (
	"context"

	"banks-api/internal/pkg/model"
	"github.com/pkg/errors"
)

func (r *BanksRepository) CreateTranslation(ctx context.Context, translation *model.Translation) (int64, error) {
	var (
		updateTranslationColums = []string{
			translationLexemeColumn,
			translationTranslatedLexemeColumn,
			translationSourceLanguageColumn,
			translationTargetLanguageColumn,
		}

		updateTranslationValues = []interface{}{
			translation.Lexeme,
			translation.TranslatedLexeme,
			translation.SourceLanguage,
			translation.TargetLanguage,
		}
	)

	if translation.ID > 0 {
		updateTranslationColums = append(updateTranslationColums, translationIDColumn)
		updateTranslationValues = append(updateTranslationValues, translation.ID)
	}

	query := psql.
		Insert(translationsTableName).
		Columns(updateTranslationColums...).
		Values(updateTranslationValues...).
		Suffix(`ON CONFLICT(id) DO UPDATE SET 
					lexeme=excluded.lexeme,
					translated_lexeme=excluded.translated_lexeme,
					source_language=excluded.source_language,
					target_language=excluded.target_language
					RETURNING id;`)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "GetBanks.ToSql")
	}

	var translationID int64
	if err = r.pool.QueryRow(ctx, rawSQL, args...).Scan(&translationID); err != nil {
		return 0, errors.Wrap(err, "UpdateBank.Exec")
	}

	return translationID, nil
}
