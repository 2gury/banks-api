package google_translate_api

import (
	translate "cloud.google.com/go/translate/apiv3"
	"context"
	"fmt"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

type TranslatorClient struct {
	client *translate.TranslationClient

	projectID string
}

func NewTranslatorClient(cli *translate.TranslationClient) *TranslatorClient {
	return &TranslatorClient{
		client:    cli,
		projectID: "numeric-rig-438915-g6",
	}
}

func (c *TranslatorClient) TranslateFromRFToKZText(ctx context.Context, text string) (string, error) {
	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", c.projectID),
		SourceLanguageCode: "ru",
		TargetLanguageCode: "kk",
		MimeType:           "text/plain",
		Contents:           []string{text},
	}

	resp, err := c.client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("TranslateText: %v", err)
	}

	for _, translation := range resp.GetTranslations() {
		return translation.GetTranslatedText(), nil
	}

	return "", fmt.Errorf("no translation for this text")
}
