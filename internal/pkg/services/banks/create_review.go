package banks

import (
	"banks-api/internal/pkg/detachutil"
	"banks-api/internal/pkg/model"
	"context"
	"golang.org/x/sync/errgroup"
	"log"
)

func (w *BanksHandler) CreateReview(ctx context.Context, review *model.Review) error {
	var (
		automoderationStrategy bool

		reviewID int64
	)

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() (err error) {
		reviewID, err = w.repo.CreateReview(egCtx, review)
		return err
	})

	eg.Go(func() error {
		strat, err := w.repo.GetAutomoderationStrategy(egCtx)
		if err != nil {
			log.Printf("error when get automoderation strategy: %s", err.Error())
		}

		automoderationStrategy = strat.AutomoderationStrategy

		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	log.Println("automoderation strategy: ", automoderationStrategy)
	if automoderationStrategy {
		go w.validateReview(detachutil.Detach(ctx), &model.Review{
			ID:        reviewID,
			Content:   review.Content,
			UserEmail: review.UserEmail,
			UserPhone: review.UserPhone,
			Rating:    review.Rating,
			BankID:    review.BankID,
		})
	}

	return nil
}

func (w *BanksHandler) validateReview(ctx context.Context, review *model.Review) {
	isValidReview, err := w.chatGPTService.IsValidReview(ctx, review.Content)
	log.Println("got answer be review:", isValidReview, err)
	if err != nil {
		return
	}

	if isValidReview {
		if err = w.UpdateReview(ctx, review.ID, true); err != nil {
			return
		}
	}
}
