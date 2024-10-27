package banks

import (
	"banks-api/internal/pkg/model"
	"banks-api/pkg/banks"
	"context"
)

type BanksService interface {
	GetBanks(ctx context.Context, filters *model.BankFilters) ([]*model.Bank, error)
	UpdateBank(ctx context.Context, bank *model.Bank) error
	DeleteBank(ctx context.Context, id int64) error
	GetPossibleBanks(_ context.Context) (*model.OffersResponse, error)
	RequestBankInformation(ctx context.Context, externalID int64) (*model.Bank, error)

	RequestTranslationText(ctx context.Context, language, text string) (string, error)
	GetTranslationByText(ctx context.Context, language, text string) (*model.Translation, error)
	CreateTranslation(ctx context.Context, translation *model.Translation) (*model.Translation, error)
	GetTranslations(ctx context.Context) ([]*model.Translation, error)

	GetReviews(ctx context.Context) ([]*model.Review, error)
	CreateReview(ctx context.Context, review *model.Review) error
	UpdateAutomoderationStrategy(ctx context.Context, automoderationEnable bool) error
	UpdateReview(ctx context.Context, reviewID int64, isApproved bool) error
}

type BanksHandler struct {
	banks.UnimplementedBanksServer

	service BanksService
}

func NewBanksHandler(service BanksService) *BanksHandler {
	return &BanksHandler{
		service: service,
	}
}
