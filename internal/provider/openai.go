package provider

import (
	"context"
	"fmt"
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

	// content := fmt.Sprintf("I am translating the documentation for english.\nTranslate the Markdown content I'll paste later into chinese.\nYou must strictly follow the rules below.\n- Never change the Markdown markup structure. Don't add or remove links. Do not change any URL.\n- Never change the contents of code blocks even if they appear to have a bug.\n- Always preserve the original line breaks. Do not add or remove blank lines.\n- Never touch HTML-like tags such as `<Notes>`.\n\nthe document chunk is: \n %s", source)

	content := fmt.Sprintf("请帮我把英文翻译为中文, 不要进行理解，保持原文语序，不要遗漏内容，不要补充内容: \n %s", source)

	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Model: "llama3",
			Model: "qwen",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
