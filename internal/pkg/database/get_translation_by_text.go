package database

import (
	"banks-api/internal/pkg/database/schema"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetTranslationByText(ctx context.Context, language, text string) (*schema.Translation, error) {
	query := psql.
		Select(
			translationIDColumn,
			translationLexemeColumn,
			translationTranslatedLexemeColumn,
			translationSourceLanguageColumn,
			translationTargetLanguageColumn,
		).From(translationsTableName).
		Where(
			sq.And{
				sq.Eq{
					translationTargetLanguageColumn: language,
				},
				sq.Eq{
					translationLexemeColumn: text,
				},
			})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBanks.ToSql")
	}

	var res []schema.Translation
	err = pgxscan.Select(ctx, r.pool, &res, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "GetTranslationByText.Get")
	}

	if len(res) > 0 {
		return &res[0], nil
	}

	return nil, nil
}
