package provider

import (
	"fmt"
	"os"

	"github.com/117503445/markdown-translate/pkg/model"
)

func GetProvider(providerCfg map[string]string) (model.Provider, error) {
	providerType, ok := providerCfg["type"]
	if !ok {
		return nil, fmt.Errorf("provider type not found")
	}

	// delete(providerCfg, "type")
	copiedCfg := make(map[string]string)
	for k, v := range providerCfg {
		if k != "type" {
			copiedCfg[k] = v
		}
	}

	switch providerType {
	case "google":
		return NewGoogleProvider(), nil
	case "mock":
		return NewMockProvider(), nil
	case "libre":
		return NewLibreProvider(), nil
	case "openai":
		return NewOpenAIProvider(providerCfg), nil
	case "uni":
		cfg := &UniProviderConfig{
			Platform: os.Getenv("UNI_PLATFORM"),
			Address:  os.Getenv("UNI_ADDRESS"),
			Key:      os.Getenv("UNI_KEY"),
		}
		return NewUniProvider(cfg), nil
	default:
		return nil, fmt.Errorf("provider %s not found", providerType)
	}
}
