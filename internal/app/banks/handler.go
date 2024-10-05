package banks

import (
	"context"

	"banks-api/internal/pkg/model"
	"banks-api/pkg/banks"
)

type BanksService interface {
	GetBanks(ctx context.Context) ([]*model.Bank, error)
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
