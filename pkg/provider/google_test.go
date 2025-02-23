package provider_test

import (
	"testing"

	"github.com/117503445/markdown-translate/pkg/provider"

	"github.com/stretchr/testify/assert"
)

func TestGoogleProvider_Translate(t *testing.T) {
	assert := assert.New(t)

	p := provider.NewGoogleProvider()

	text, err := p.Translate("hello")
	assert.Nil(err)

	t.Log(text)
}
