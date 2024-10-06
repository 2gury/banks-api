package banks

import (
	"context"

	"banks-api/internal/pkg/database/schema"
	"banks-api/internal/pkg/model"
)

type BanksRepository interface {
	GetBanks(ctx context.Context, filters *model.BankFilters) ([]schema.Bank, error)
}

type BanksHandler struct {
	repo BanksRepository
}

func NewBanksService(repo BanksRepository) *BanksHandler {
	return &BanksHandler{
		repo: repo,
	}
}
