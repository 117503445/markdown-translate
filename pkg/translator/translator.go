package translator

import (
	"bytes"

	"strings"

	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Translator struct {
	provider provider.Provider
}

func NewTranslator(translateProvider provider.Provider) *Translator {
	p := translateProvider
	if p == nil {
		p = provider.NewGoogleProvider()
	}
	return &Translator{provider: p}
}

func (t *Translator) Translate(source string) (string, error) {
	// return t.provider.Translate(source)
	var buf bytes.Buffer

	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	src := []byte(source)

	doc := markdown.Parser().Parse(text.NewReader(src))

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n := node.(type) {
			case *ast.Heading:
				level := n.Level
				buf.WriteString(strings.Repeat("#", level) + " ")
				buf.WriteString(string(n.Text(src)))
				buf.WriteString(" ")
				translated, err :=

					t.provider.Translate(string(n.Text(src)))
				if err != nil {
					return ast.WalkStop, err
				}
				buf.WriteString(translated)
				buf.WriteString("\n")
				return ast.WalkSkipChildren, nil
			case *ast.Paragraph:
				buf.WriteString("\n")
			case *ast.Text:
				buf.WriteString(string(n.Text(src)))
				buf.WriteString("\n\n")
				translated, err := t.provider.Translate(string(n.Text(src)))
				if err != nil {
					return ast.WalkStop, err
				}
				buf.WriteString(translated)
				buf.WriteString("\n")
			}
		}
		return ast.WalkContinue, nil
	})

	return buf.String(), nil
}
