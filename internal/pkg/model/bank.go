package model

import "banks-api/internal/pkg/database/schema"

type Bank struct {
	ID       int64
	BankName string
}

func ConvertBanksToModel(banks []schema.Bank) []*Bank {
	modelBanks := make([]*Bank, 0, len(banks))

	for _, b := range banks {
		modelBanks = append(modelBanks, ConvertBankToModel(b))
	}

	return modelBanks
}

func ConvertBankToModel(bank schema.Bank) *Bank {
	return &Bank{
		ID:       bank.ID,
		BankName: bank.BankName,
	}
}
