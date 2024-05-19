// https://github.com/xgd16/UniTranslate

package provider

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
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
	p.client.R().SetHeader("Content-Type", "application/json").SetBody(map[string]string{
		"from":     "auto",
		"to":       "en",
		"text":     source,
		"platform": p.platform,
	})
	log.Debug().Str("source", source).Msg("translating")
	return fmt.Sprintf("[翻译]%s[结束]", source), nil
}
