package banks

import (
	"context"

	"banks-api/internal/app/banks/mappers"
	"banks-api/internal/pkg/model"
	pbbanks "banks-api/pkg/banks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *BanksHandler) GetBanks(ctx context.Context, req *pbbanks.GetBanksRequest) (*pbbanks.GetBanksResponse, error) {
	banks, err := w.service.GetBanks(ctx, &model.BankFilters{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbbanks.GetBanksResponse{
		Banks: mappers.ModelToPBBanks(banks),
	}, nil
}
