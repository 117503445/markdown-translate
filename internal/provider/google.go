package provider

import (
	"log"

	gtranslate "github.com/gilang-as/google-translate"
)

type GoogleProvider struct {
}

func NewGoogleProvider() *GoogleProvider {
	return &GoogleProvider{}
}

func (p *GoogleProvider) Translate(source string) (string, error) {
	log.Printf("< %s\n", source)
	value := gtranslate.Translate{
		Text: source,
		From: "en",
		To:   "zh",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		return "", err
	} else {
		text := translated.Text
		log.Printf("> %s\n", text)
		return text, nil
	}
}
