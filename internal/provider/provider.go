package provider

import (
	"fmt"
	"os"

	"github.com/117503445/markdown-translate/pkg/model"
)

func GetProvider(provider string) (model.Provider, error) {
	switch provider {
	case "google":
		return NewGoogleProvider(), nil
	case "mock":
		return NewMockProvider(), nil
	case "libre":
		return NewLibreProvider(), nil
	case "openai":
		return NewOpenAIProvider(), nil
	case "uni":
		cfg := &UniProviderConfig{
			Platform: os.Getenv("UNI_PLATFORM"),
			Address:  os.Getenv("UNI_ADDRESS"),
			Key:      os.Getenv("UNI_KEY"),
		}
		return NewUniProvider(cfg), nil
	default:
		return nil, fmt.Errorf("provider %s not found", provider)
	}
}
