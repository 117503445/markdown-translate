package provider

import (
	tr "github.com/snakesel/libretranslate"
)

// https://github.com/LibreTranslate/LibreTranslate

type LibreProvider struct {
	client *tr.Translation
}

func NewLibreProvider() *LibreProvider {
	client := tr.New(tr.Config{
		Url: "https://libretranslate.com",
		Key: "",
	})
	return &LibreProvider{client: client}
}

func (p *LibreProvider) Translate(source string) (string, error) {
	return p.client.Translate(source, "auto", "zh")
}
