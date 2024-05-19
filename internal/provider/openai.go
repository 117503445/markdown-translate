package provider

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
	client *openai.Client
}

func NewOpenAIProvider() *OpenAIProvider {
	config := openai.DefaultConfig("your token")
	config.BaseURL = os.Getenv("PROVIDER_OPENAI_ENDPOINT")
	log.Debug().Msgf("OpenAI endpoint: %s", config.BaseURL)

	client := openai.NewClientWithConfig(config)
	return &OpenAIProvider{client: client}
}

func (p *OpenAIProvider) Translate(source string) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Model: "llama3",
			Model: "qwen",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Please help me to translate markdown into Chinese and keep the formatting",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: source,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
