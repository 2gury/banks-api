package banks

import (
	"banks-api/internal/pkg/database/schema"
	"banks-api/internal/pkg/model"
	"context"
	"golang.org/x/sync/errgroup"
)

func (w *BanksHandler) GetReviews(ctx context.Context) ([]*model.Review, error) {
	var (
		reviews []schema.Review

		banks []schema.Bank
	)

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() (err error) {
		reviews, err = w.repo.GetReviews(egCtx)
		return err
	})

	eg.Go(func() (err error) {
		banks, err = w.repo.GetBanks(egCtx, &model.BankFilters{})
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return model.ConvertReviewsToModel(reviews, getBanksMap(banks)), nil
}

func getBanksMap(banks []schema.Bank) map[int64]string {
	banksMap := make(map[int64]string, len(banks))

	for _, bank := range banks {
		banksMap[bank.ID] = bank.Name
	}

	return banksMap
}
