package test

import (
	"testing"

	"github.com/117503445/markdown-translate/pkg/translator"
)

func TestExample(t *testing.T) {
	source := "# Hello\n\nThis is a markdown file."
	expected := "# Hello\n\nThis is a markdown file."

	translator := translator.NewTranslator(nil)
	actual, _ := translator.Translate(source)

	if actual != expected {
		t.Errorf("Expected: %s\nActual: %s", expected, actual)
	}
}
