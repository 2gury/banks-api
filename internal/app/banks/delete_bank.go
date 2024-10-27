package banks

import (
	"context"

	pbbanks "banks-api/pkg/banks"
)

func (w *BanksHandler) DeleteBank(ctx context.Context, req *pbbanks.DeleteBankRequest) (*pbbanks.DeleteBankResponse, error) {
	if err := w.service.DeleteBank(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &pbbanks.DeleteBankResponse{}, nil
}
