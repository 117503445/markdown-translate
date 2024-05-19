package translator

import (
	"bytes"
	"strings"

	// "strings"

	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Provider interface {
	Translate(source string) (string, error)
}

type Translator struct {
	provider Provider
}

func NewTranslator(translateProvider Provider) *Translator {
	p := translateProvider
	if p == nil {
		p = provider.NewGoogleProvider()
	}
	return &Translator{provider: p}
}

func getRawText(node ast.Node, src []byte) string {
	rawText := ""
	for i := 0; i < node.Lines().Len(); i++ {
		line := node.Lines().At(i)
		rawText += string(line.Value(src))
	}
	return rawText
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

	// log.Debug().Str("Text", string(doc.Dump(src, 2))).Msg("ast.Document")
	// doc.Dump(src, 2)

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			// log.Debug().Str("Type", node.Kind().String()).Msg("Node")
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
				buf.WriteString("\n\n")
			case *ast.ThematicBreak:
				buf.WriteString("---\n\n")
			case *ast.Paragraph:
				raw := getRawText(n, src)
				_, err := buf.WriteString(raw + "\n\n")
				if err != nil {
					return ast.WalkStop, err
				}

				translated, err := t.provider.Translate(raw)
				if err != nil {
					return ast.WalkStop, err
				}

				_, err = buf.WriteString(translated + "\n\n")
				if err != nil {
					return ast.WalkStop, err
				}

			case *ast.Document:
				return ast.WalkContinue, nil
			default:
				log.Debug().Str("Type", node.Kind().String()).Str("Text", string(node.Text(src))).Str("Raw", getRawText(node, src)).
					Msg("ast.Node [ignored]")
			}
		}
		return ast.WalkSkipChildren, nil
	})

	return buf.String(), nil
}
