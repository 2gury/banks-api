package leadgid

import (
	"banks-api/internal/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	offersURL = "https://api.leadgid.com/offers/v1/affiliates/offers"

	offerInformationURL = "https://api.leadgid.com/offers/v1/affiliates/offers/"
)

type LeadGIDClient struct {
	apiKey string
}

func NewLeadGIDClient(key string) *LeadGIDClient {
	return &LeadGIDClient{
		apiKey: key,
	}
}

func (lg *LeadGIDClient) GetPossibleBanks() (*model.OffersResponse, error) {
	query := url.Values{}

	query.Add("lang", "ru")
	query.Add("department", "3")
	query.Add("product", "1")
	query.Add("convert-currency", "RUB")
	query.Add("limit", "100")

	reqURL := fmt.Sprintf("%s?%s", offersURL, query.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-account-token", lg.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var offersResponse model.OffersResponse
	if err := json.Unmarshal(body, &offersResponse); err != nil {
		return nil, err
	}

	return &offersResponse, nil
}

func (lg *LeadGIDClient) GetBankInformationByID(bankID int64) (*model.OfferInformation, error) {
	query := url.Values{}

	query.Add("lang", "ru")

	reqURL := fmt.Sprintf("%s%d?%s", offerInformationURL, bankID, query.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-account-token", lg.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var offersInfo model.OfferInformation
	if err := json.Unmarshal(body, &offersInfo); err != nil {
		return nil, err
	}

	return &offersInfo, nil
}
