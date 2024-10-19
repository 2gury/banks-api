package banks

import (
	"banks-api/internal/pkg/database/schema"
	"banks-api/internal/pkg/model"
	"context"
	"log"
)

func (w *BanksHandler) GetBanks(ctx context.Context, filters *model.BankFilters) ([]*model.Bank, error) {
	schemaBanks, err := w.repo.GetBanks(ctx, filters)
	if err != nil {
		return nil, err
	}

	banks, err := model.ConvertBanksToModel(schemaBanks)
	if err != nil {
		return nil, err
	}

	if contains(filters.Language, kzSymbols) {
		translations, err := w.repo.GetTranslations(ctx)
		if err != nil {
			log.Println("error when get translations", err)
			return banks, nil
		}

		translationsMap := getMapTranslations(translations)

		for _, b := range banks {
			translateBankData(b, translationsMap)
		}

	}

	return banks, nil
}

func getMapTranslations(translations []schema.Translation) map[string]string {
	translationsMap := make(map[string]string, len(translations))

	for _, tr := range translations {
		translationsMap[tr.Lexeme] = tr.TranslatedLexeme
	}

	return translationsMap
}

func translateBankData(bank *model.Bank, translationsMap map[string]string) bool {
	var (
		bankData            = bank.BankData
		needGetTranslations = false
	)
	for id, r := range bankData.Registration {
		val, ok := translationsMap[r]
		if !ok {
			needGetTranslations = true
		} else {
			bank.BankData.Registration[id] = val
		}
	}

	for id, cp := range bankData.CreditPurpose {
		val, ok := translationsMap[cp]
		if !ok {
			needGetTranslations = true
		} else {
			bank.BankData.CreditPurpose[id] = val
		}
	}

	for id, d := range bankData.Documents {
		val, ok := translationsMap[d]
		if !ok {
			needGetTranslations = true
		} else {
			bank.BankData.Documents[id] = val
		}
	}

	for id, om := range bankData.ObtainMethod {
		val, ok := translationsMap[om]
		if !ok {
			needGetTranslations = true
		} else {
			bank.BankData.ObtainMethod[id] = val
		}
	}

	return needGetTranslations
}
