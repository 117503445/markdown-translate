package translator

import "github.com/117503445/markdown-translate/internal/provider"

type Translator struct {
	provider provider.Provider
}

func NewTranslator() *Translator {
	provider := provider.NewMockProvider()
	return &Translator{provider: provider}
}

func (t *Translator) Translate(source string) string {
	return t.provider.Translate(source)
}
