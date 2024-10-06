package mappers

import (
	"banks-api/internal/pkg/model"
	pbbanks "banks-api/pkg/banks"
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
		Id:            bank.ID,
		Name:          bank.Name,
		Logo:          bank.Logo,
		Url:           bank.URL,
		PeriodFrom:    bankData.PeriodFrom,
		PeriodTo:      bankData.PeriodTo,
		AmountFrom:    bankData.AmountFrom,
		AmountTo:      bankData.AmountTo,
		RateFrom:      bankData.RateFrom,
		RateTo:        bankData.RateTo,
		ReviewTime:    bankData.ReviewTime,
		Registration:  bankData.Registration,
		CreditPurpose: bankData.CreditPurpose,
		Documents:     bankData.Documents,
		ObtainMethod:  bankData.ObtainMethod,
	}
}
