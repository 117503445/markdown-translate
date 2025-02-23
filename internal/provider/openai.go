package provider

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
	client *openai.Client

	model string
}

func MustGet(m map[string]string, key string) string {
	v, ok := m[key]
	if !ok {
		log.Fatal().Interface("map", m).Str("key", key).Msg("provider config not found")
	}
	return v
}

func OptionalGet(m map[string]string, key string, defaultValue string) string {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}
	return v
}

func NewOpenAIProvider(providerCfg map[string]string) *OpenAIProvider {
	token := OptionalGet(providerCfg, "token", "")
	baseUrl := OptionalGet(providerCfg, "host", "")
	model := OptionalGet(providerCfg, "model", "")

	config := openai.DefaultConfig(token)
	config.BaseURL = baseUrl

	client := openai.NewClientWithConfig(config)
	return &OpenAIProvider{client: client, model: model}
}

func (p *OpenAIProvider) Translate(source string) (string, error) {
	content := fmt.Sprintf("请将以下提供的英文Markdown文本翻译成中文。请注意，翻译过程中需严格保持原文语序，不得对原文内容做任何解释或补充。确保所有内容都被翻译，没有遗漏。英文Markdown文本如下:\n %s", source)
	log.Debug().Str("content", content).Send()

	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: p.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("openai translate failed")
		return "", err
	}

	response := resp.Choices[0].Message.Content
	log.Debug().Str("response", response).Send()

	return response, nil
}
