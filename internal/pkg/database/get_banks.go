package database

import (
	"context"
	"fmt"

	"banks-api/internal/pkg/database/schema"
	"banks-api/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *BanksRepository) GetBanks(ctx context.Context, filters *model.BankFilters) ([]schema.Bank, error) {
	query := psql.
		Select(
			bankIDColumn,
			bankExternalIDColumn,
			bankExternalLegacyIDColumn,
			bankNameColumn,
			bankLogoColumn,
			bankURLColumn,
			bankDataColumn,
			bankCreatedAtColumn,
			bankUpdatedAtColumn,
		).From(banksTableName).
		OrderBy(fmt.Sprintf("%s DESC", bankIDColumn))

	query = applyBanksFilter(query, filters)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBanks.ToSql")
	}

	var res []schema.Bank
	err = pgxscan.Select(ctx, r.pool, &res, rawSQL, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "GetBanks.Select")
	}

	return res, nil
}

func applyBanksFilter(qb sq.SelectBuilder, filter *model.BankFilters) sq.SelectBuilder {
	if filter.Limit != 0 {
		qb = qb.Limit(filter.Limit)
	}

	if filter.Offset != 0 {
		qb = qb.Offset(filter.Offset)
	}

	return qb
}
