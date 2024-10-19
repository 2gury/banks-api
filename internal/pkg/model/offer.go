package model

type OffersResponse struct {
	Data []OfferData `json:"data"`
}

type OfferData struct {
	Id       int    `json:"id"`
	LegacyId int    `json:"legacy_id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
}

type OfferInformation struct {
	Data OfferFullData `json:"data"`
}

type OfferFullData struct {
	Id          int    `json:"id"`
	LegacyId    int    `json:"legacy_id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	DefaultUrl  string `json:"default_url"`
}
