// https://github.com/xgd16/UniTranslate

package provider

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

type UniProvider struct {
	client   *resty.Client
	platform string
	address  string
	key      string
}

type UniProviderConfig struct {
	Platform string
	Address  string
	Key      string
}

func NewUniProvider(cfg *UniProviderConfig) *UniProvider {
	client := resty.New()
	return &UniProvider{client: client, platform: cfg.Platform, address: cfg.Address, key: cfg.Key}
}

func (p *UniProvider) Translate(source string) (string, error) {
	resp, err := p.client.R().SetHeader("Content-Type", "application/json").SetBody(map[string]string{
		"from":     "auto",
		"to":       "zh-CHS",
		"text":     source,
		"platform": p.platform,
	}).Post(p.address + "?key=" + p.key)
	if err != nil {
		return "", err
	}

	// log.Debug().Str("resp", resp.String()).Msg("response")

	text := gjson.Get(resp.String(), "data.translate.0.text").String()

	log.Debug().Str("source", source).Str("text", text).Msg("translating")
	return text, nil
}
