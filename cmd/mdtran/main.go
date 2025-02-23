package main

import (
	"os"
	"time"

	"github.com/117503445/goutils"
	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/117503445/markdown-translate/pkg/cfg"
	"github.com/117503445/markdown-translate/pkg/translator"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()
	cfg.Load("config.toml")

	inputText, err := os.ReadFile(cfg.Cfg.Target)
	if err != nil {
		log.Fatal().Err(err).Msg("read input file")
	}

	// first try to use input file name + .out.md
	// if exists, use input file name + date + .out.md
	// if still exists, error
	var output string
	if _, err := os.Stat(cfg.Cfg.Target + ".out.md"); os.IsNotExist(err) {
		output = cfg.Cfg.Target + ".out.md"
	} else {
		date := time.Now().Format("20060102-150405")
		output = cfg.Cfg.Target + "." + date + ".out.md"
		if _, err := os.Stat(output); err == nil {
			log.Fatal().Strs("output", []string{cfg.Cfg.Target + ".out.md", output}).Msg("output file exists")
		}
	}

	provider, err := provider.GetProvider(cfg.Cfg.Provider)
	if err != nil {
		log.Fatal().Err(err).Msg("provider not found")
	}

	translator := translator.NewTranslator(provider)

	outputText, err := translator.Translate(string(inputText))
	if err != nil {
		log.Fatal().Err(err).Msg("translate failed")
	}

	err = os.WriteFile(output, []byte(outputText), 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("write output file")
	}
}
