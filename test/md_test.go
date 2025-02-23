package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/117503445/goutils"
	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/117503445/markdown-translate/internal/provider/cache"
	"github.com/117503445/markdown-translate/pkg/cfg"
	"github.com/117503445/markdown-translate/pkg/translator"
	"github.com/117503445/markdown-translate/test/examples"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	goutils.InitZeroLog()

	args := os.Args
	os.Args = os.Args[:1]
	// log.Debug().Strs("args", os.Args).Send()
	cfg.Load("/workspace/config.toml")
	os.Args = args

	// TODO load level from cfg
	log.Logger = log.Level(zerolog.TraceLevel)

	m.Run()
}

func TestMockAll(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslator(provider.NewMockProvider())

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)
		assert.Nil(err)

		err = goutils.WriteText(fmt.Sprintf("./examples/%v.mock.out", k), r)
		assert.Nil(err)
	}
	log.Info().Msg("mock all done")
}

func TestCfg(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslatorByCfg()

	for k, v := range examples.Examples {
		r, err := translator.Translate(v)
		assert.Nil(err)

		err = goutils.WriteText(fmt.Sprintf("./examples/%v.cfg.out", k), r)
		assert.Nil(err)
	}
	log.Info().Msg("mock all done")
}

func TestOpenAI(t *testing.T) {
	assert := assert.New(t)

	translator := translator.NewTranslator(provider.NewOpenAIProvider(map[string]string{}))

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
