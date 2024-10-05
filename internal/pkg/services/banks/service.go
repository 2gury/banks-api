package banks

import (
	"context"

	"banks-api/internal/pkg/database/schema"
)

type BanksRepository interface {
	GetBanks(ctx context.Context) ([]schema.Bank, error)
}

type BanksHandler struct {
	repo BanksRepository
}

func NewBanksService(repo BanksRepository) *BanksHandler {
	return &BanksHandler{
		repo: repo,
	}
}
