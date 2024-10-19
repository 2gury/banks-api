package banks

import (
	"banks-api/internal/pkg/database/schema"
	"banks-api/internal/pkg/model"
	"context"
)

var (
	kzSymbols = []string{
		"kz", "kk", "Казахский", "Казахстан", "КЗ", "кз", "Kazakh", "kazakhstan",
	}
)

type BanksRepository interface {
	GetBanks(ctx context.Context, filters *model.BankFilters) ([]schema.Bank, error)
	UpdateBank(ctx context.Context, bank *model.Bank) (int64, error)

	GetTranslationByText(ctx context.Context, language, text string) (*schema.Translation, error)
	CreateTranslation(ctx context.Context, translation *model.Translation) (int64, error)
	GetTranslations(ctx context.Context) ([]schema.Translation, error)
}

type TranslatorService interface {
	TranslateFromRFToKZText(ctx context.Context, text string) (string, error)
}

type LeadgidService interface {
	GetPossibleBanks() (*model.OffersResponse, error)
	GetBankInformationByID(bankID int64) (*model.OfferInformation, error)
}

type ChatGPTService interface {
	GrabText(ctx context.Context, text string) (*model.BankChatGPTInfo, error)
}

type BanksHandler struct {
	repo              BanksRepository
	translatorService TranslatorService
	leadgidService    LeadgidService
	chatGPTService    ChatGPTService
}

func NewBanksService(
	repo BanksRepository,
	translatorService TranslatorService,
	leadgidService LeadgidService,
	chatGPTService ChatGPTService,
) *BanksHandler {
	return &BanksHandler{
		repo:              repo,
		translatorService: translatorService,
		leadgidService:    leadgidService,
		chatGPTService:    chatGPTService,
	}
}

func contains(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
