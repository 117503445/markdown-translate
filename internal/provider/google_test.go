package provider_test

import (
	"github.com/117503445/markdown-translate/internal/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoogleProvider_Translate(t *testing.T) {
	assert := assert.New(t)

	p := provider.NewGoogleProvider()

	text, err := p.Translate("hello")
	assert.Nil(err)

	t.Log(text)
}
