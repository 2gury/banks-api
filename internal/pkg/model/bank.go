package model

import (
	"encoding/json"
	"fmt"
	"time"

	"banks-api/internal/pkg/database/schema"
)

type Bank struct {
	ID               int64
	ExternalID       int64
	ExternalLegacyID int64
	Name             string
	Logo             string
	URL              string

	BankData BankData

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BankData struct {
	PeriodFrom    int64    `json:"period_from"`
	PeriodTo      int64    `json:"period_to"`
	AmountFrom    int64    `json:"amount_from"`
	AmountTo      int64    `json:"amount_to"`
	RateFrom      float64  `json:"rate_from"`
	RateTo        float64  `json:"rate_to"`
	Registration  []string `json:"registration"`
	CreditPurpose []string `json:"credit_purpose"`
	Documents     []string `json:"documents"`
	ReviewTime    int64    `json:"review_time"`
	ObtainMethod  []string `json:"obtain_method"`
	Description   string   `json:"description"`
}

func ConvertBanksToModel(banks []schema.Bank) ([]*Bank, error) {
	if banks == nil {
		return nil, fmt.Errorf("empty banks")
	}

	modelBanks := make([]*Bank, 0, len(banks))

	for _, b := range banks {
		modelBank, err := ConvertBankToModel(b)
		if err != nil {
			return nil, err
		}

		modelBanks = append(modelBanks, modelBank)
	}

	return modelBanks, nil
}

func ConvertBankToModel(bank schema.Bank) (*Bank, error) {
	var bankData BankData
	if err := json.Unmarshal([]byte(bank.BankData), &bankData); err != nil {
		return nil, err
	}

	return &Bank{
		ID:               bank.ID,
		ExternalID:       bank.ExternalID,
		ExternalLegacyID: bank.ExternalLegacyID,
		Name:             bank.Name,
		Logo:             bank.Logo,
		URL:              bank.URL,
		BankData:         bankData,
		CreatedAt:        bank.CreatedAt,
		UpdatedAt:        bank.UpdatedAt,
	}, nil
}

func ConvertBankToSchema(bank *Bank) (*schema.Bank, error) {
	if bank == nil {
		return nil, fmt.Errorf("empty bank")
	}

	rawDataBank, err := json.Marshal(bank.BankData)
	if err != nil {
		return nil, err
	}

	return &schema.Bank{
		ID:               bank.ID,
		ExternalID:       bank.ExternalID,
		ExternalLegacyID: bank.ExternalLegacyID,
		Name:             bank.Name,
		Logo:             bank.Logo,
		URL:              bank.URL,
		BankData:         string(rawDataBank),
		CreatedAt:        bank.CreatedAt,
		UpdatedAt:        bank.UpdatedAt,
	}, nil
}

type BankChatGPTInfo struct {
	CompanyName   string   `json:"company_name"`
	CompanyLogo   string   `json:"company_logo"`
	PeriodFrom    int      `json:"period_from"`
	PeriodTo      int      `json:"period_to"`
	AmountFrom    int      `json:"amount_from"`
	AmountTo      int      `json:"amount_to"`
	RateFrom      float64  `json:"rate_from"`
	RateTo        float64  `json:"rate_to"`
	Registration  []string `json:"registration"`
	CreditPurpose []string `json:"credit_purpose"`
	Documents     []string `json:"documents"`
	ReviewTime    int      `json:"review_time"`
	ObtainMethod  []string `json:"obtain_method"`
	Description   string   `json:"description"`
}
