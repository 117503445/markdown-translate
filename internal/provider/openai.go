package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
}

func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{}
}

func (p *OpenAIProvider) Translate(source string) (string, error) {
	config := openai.DefaultConfig("your token")
	config.BaseURL = os.Getenv("PROVIDER_OPENAI_ENDPOINT")
	log.Debug().Msgf("OpenAI endpoint: %s", config.BaseURL)

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
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
					Content: "Prior work has focused on increasing LBFT performance by improving the commit phase, e.g., reducing message complex-ity [19], truncating communication rounds [26], and enhancing tolerance to Byzantine faults [27], [28]. Recent works [23], [25] reveal that a more significant factor limiting LBFT's scalability lies in the proposing phase, in which a proposal with batched transaction data (e.g., $10\\mathrm{MB}$ ) is disseminated by the single leader node, whereas messages exchanged in the commit phase (e.g., signatures, hashes) are much smaller (e.g., 100 Byte). Formal analysis in Appendix A shows that reducing the message complexity of the commit phase cannot address this scalability issue.",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
