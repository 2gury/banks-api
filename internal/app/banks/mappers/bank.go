package mappers

import (
	"banks-api/internal/pkg/model"
	pbbanks "banks-api/pkg/banks"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelToPBBanks(banks []*model.Bank) []*pbbanks.Bank {
	pbBanks := make([]*pbbanks.Bank, 0, len(banks))

	for _, b := range banks {
		pbBanks = append(pbBanks, ModelToPBBank(b))
	}

	return pbBanks
}

func ModelToPBBank(bank *model.Bank) *pbbanks.Bank {
	bankData := bank.BankData
	return &pbbanks.Bank{
		Id:               bank.ID,
		ExternalId:       bank.ExternalID,
		ExternalLegacyId: bank.ExternalLegacyID,
		Name:             bank.Name,
		Logo:             bank.Logo,
		Url:              bank.URL,
		CreatedAt:        timestamppb.New(bank.CreatedAt),
		UpdatedAt:        timestamppb.New(bank.UpdatedAt),
		PeriodFrom:       bankData.PeriodFrom,
		PeriodTo:         bankData.PeriodTo,
		AmountFrom:       bankData.AmountFrom,
		AmountTo:         bankData.AmountTo,
		RateFrom:         bankData.RateFrom,
		RateTo:           bankData.RateTo,
		ReviewTime:       bankData.ReviewTime,
		Registration:     bankData.Registration,
		CreditPurpose:    bankData.CreditPurpose,
		Documents:        bankData.Documents,
		ObtainMethod:     bankData.ObtainMethod,
	}
}

func PBToModelBank(bank *pbbanks.Bank) *model.Bank {
	modelBank := &model.Bank{
		ID:               bank.Id,
		ExternalID:       bank.ExternalId,
		ExternalLegacyID: bank.ExternalLegacyId,
		Name:             bank.Name,
		Logo:             bank.Logo,
		URL:              bank.Url,
		CreatedAt:        bank.CreatedAt.AsTime(),
		UpdatedAt:        bank.UpdatedAt.AsTime(),
		BankData: model.BankData{
			PeriodFrom:    bank.PeriodFrom,
			PeriodTo:      bank.PeriodTo,
			AmountFrom:    bank.AmountFrom,
			AmountTo:      bank.AmountTo,
			RateFrom:      bank.RateFrom,
			RateTo:        bank.RateTo,
			ReviewTime:    bank.ReviewTime,
			Registration:  bank.Registration,
			CreditPurpose: bank.CreditPurpose,
			Documents:     bank.Documents,
			ObtainMethod:  bank.ObtainMethod,
		},
	}

	if bank.CreatedAt != nil {
		modelBank.CreatedAt = bank.CreatedAt.AsTime()
	}

	if bank.UpdatedAt != nil {
		modelBank.UpdatedAt = bank.UpdatedAt.AsTime()
	}

	return modelBank
}
