package provider

import (
	gtranslate "github.com/gilang-as/google-translate"
)

type GoogleProvider struct {
}

func NewGoogleProvider() *GoogleProvider {
	return &GoogleProvider{}
}

func (p *GoogleProvider) Translate(source string) string {
	value := gtranslate.Translate{
		Text: source,
		From: "en",
		To:   "zh",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	} else {
		return translated.Text
	}
}
