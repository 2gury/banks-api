package banks

import (
	"context"

	"banks-api/internal/pkg/model"
	"banks-api/pkg/banks"
)

type BanksService interface {
	GetBanks(ctx context.Context, filters *model.BankFilters) ([]*model.Bank, error)
	UpdateBank(ctx context.Context, bank *model.Bank) error
	GetPossibleBanks(_ context.Context) (*model.OffersResponse, error)
	RequestBankInformation(ctx context.Context, externalID int64) (*model.Bank, error)

	RequestTranslationText(ctx context.Context, language, text string) (string, error)
	GetTranslationByText(ctx context.Context, language, text string) (*model.Translation, error)
	CreateTranslation(ctx context.Context, translation *model.Translation) (*model.Translation, error)
	GetTranslations(ctx context.Context) ([]*model.Translation, error)
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
