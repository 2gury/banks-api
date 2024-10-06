package banks

import (
	"context"

	"banks-api/internal/app/banks/mappers"
	pbbanks "banks-api/pkg/banks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *BanksHandler) UpdateBank(ctx context.Context, req *pbbanks.UpdateBankRequest) (*pbbanks.UpdateBankResponse, error) {
	if err := validateUpdateBankRequest(req); err != nil {
		return nil, err
	}

	err := w.service.UpdateBank(ctx, mappers.PBToModelBank(req.GetBank()))
	if err != nil {
		return nil, err
	}

	return &pbbanks.UpdateBankResponse{}, nil
}

func validateUpdateBankRequest(req *pbbanks.UpdateBankRequest) error {
	if req.GetBank() == nil {
		return status.Error(codes.InvalidArgument, "empty bank")
	}

	bank := req.Bank
	if bank.GetName() == "" || bank.GetLogo() == "" || bank.GetUrl() == "" {
		return status.Error(codes.InvalidArgument, "required field is empty ")
	}

	return nil
}
