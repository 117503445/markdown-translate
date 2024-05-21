package test

import (
	"os"
	"testing"

	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/117503445/markdown-translate/internal/provider/cache"
	"github.com/117503445/markdown-translate/pkg/translator"
	"github.com/117503445/markdown-translate/test/examples"
	"github.com/stretchr/testify/assert"
)

func TestMockAll(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslator(provider.NewMockProvider())

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)

		assert.Nil(err)

		os.WriteFile("./examples/"+k+".mock.out", []byte(r), 0644)
	}
}

func TestOpenAI(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslator(provider.NewOpenAIProvider())

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)

		assert.Nil(err)

		os.WriteFile("./examples/"+k+".openai.out", []byte(r), 0644)
	}
}

func TestGoogleAll(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslator(provider.NewGoogleProvider())

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)

		assert.Nil(err)

		os.WriteFile("./examples/"+k+".google.out", []byte(r), 0644)
	}
}

func TestUniAll(t *testing.T) {
	assert := assert.New(t)

	c := cache.NewBadgerWithConfig(&cache.BadgerConfig{
		Dir: "./data/uni",
	})

	p := provider.NewUniProvider(&provider.UniProviderConfig{
		Platform: "",
		Address:  "http://192.168.100.226:9431/api/translate",
		Key:      "hdasdhasdhsahdkasjfsoufoqjoje",
	})

	translator := translator.NewTranslatorWithConfig(
		&translator.TranslatorConfig{
			Provider: p,
			Cache:    c,
		},
	)

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)

		assert.Nil(err)

		os.WriteFile("./examples/"+k+".uni.out", []byte(r), 0644)
	}
}
