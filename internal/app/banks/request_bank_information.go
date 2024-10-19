package banks

import (
	"banks-api/internal/app/banks/mappers"
	"banks-api/pkg/banks"
	"context"
)

func (w *BanksHandler) RequestBankInformation(ctx context.Context, req *banks.RequestBankInformationRequest) (*banks.RequestBankInformationResponse, error) {
	bankInfo, err := w.service.RequestBankInformation(ctx, req.GetExternalId())
	if err != nil {
		return nil, err
	}

	return &banks.RequestBankInformationResponse{
		Bank: mappers.ModelToPBBank(bankInfo),
	}, nil
}
