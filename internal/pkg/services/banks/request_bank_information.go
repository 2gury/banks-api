package banks

import (
	"banks-api/internal/pkg/model"
	"context"
	"time"
)

func (w *BanksHandler) RequestBankInformation(ctx context.Context, externalID int64) (*model.Bank, error) {
	offerInfo, err := w.leadgidService.GetBankInformationByID(externalID)
	if err != nil {
		return nil, err
	}

	chatGPTInfo, err := w.chatGPTService.GrabText(ctx, offerInfo.Data.Description)
	if err != nil {
		return nil, err
	}

	return &model.Bank{
		ExternalID:       int64(offerInfo.Data.Id),
		ExternalLegacyID: int64(offerInfo.Data.LegacyId),
		Name:             offerInfo.Data.Name,
		Logo:             offerInfo.Data.Logo,
		URL:              offerInfo.Data.DefaultUrl,
		UpdatedAt:        time.Now(),
		CreatedAt:        time.Now(),
		BankData: model.BankData{
			PeriodFrom:    int64(chatGPTInfo.PeriodFrom),
			PeriodTo:      int64(chatGPTInfo.PeriodTo),
			AmountFrom:    int64(chatGPTInfo.AmountFrom),
			AmountTo:      int64(chatGPTInfo.AmountTo),
			RateFrom:      chatGPTInfo.RateFrom,
			RateTo:        chatGPTInfo.RateTo,
			Registration:  chatGPTInfo.Registration,
			CreditPurpose: chatGPTInfo.CreditPurpose,
			Documents:     chatGPTInfo.Documents,
			ReviewTime:    int64(chatGPTInfo.ReviewTime),
			ObtainMethod:  chatGPTInfo.ObtainMethod,
			Description:   chatGPTInfo.Description,
		},
	}, nil
}
