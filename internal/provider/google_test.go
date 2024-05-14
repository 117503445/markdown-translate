package provider_test

import (
	"github.com/117503445/markdown-translate/internal/provider"
	"testing"
)

func TestGoogleProvider_Translate(t *testing.T) {
	p := provider.NewGoogleProvider()
	t.Log(p.Translate("hello"))
}
