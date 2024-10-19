package database

import (
	"context"

	"banks-api/internal/pkg/database/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetTranslations(ctx context.Context) ([]schema.Translation, error) {
	query := psql.
		Select(
			translationIDColumn,
			translationLexemeColumn,
			translationTranslatedLexemeColumn,
			translationSourceLanguageColumn,
			translationTargetLanguageColumn,
		).From(translationsTableName)

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

		return nil, errors.Wrap(err, "GetTranslations.Select")
	}

	return res, nil
}
