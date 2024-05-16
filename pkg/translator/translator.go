package translator

import (
	"bytes"

	"strings"

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
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.Heading")
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
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.Paragraph")
				buf.WriteString("\n")
			case *ast.CodeSpan:
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.CodeSpan")
				buf.WriteString("`")
				buf.WriteString(string(n.Text(src)))
				buf.WriteString("`")
				return ast.WalkSkipChildren, nil
			case *ast.Image:
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.Image")
				buf.WriteString("![")
				buf.WriteString(string(n.Text(src)))
				buf.WriteString("](")
				buf.WriteString(string(n.Destination))
				buf.WriteString(")")
				return ast.WalkSkipChildren, nil
			case *ast.Text:
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.Text")
				buf.WriteString(string(n.Text(src)))
				buf.WriteString("\n\n")
				translated, err := t.provider.Translate(string(n.Text(src)))
				if err != nil {
					return ast.WalkStop, err
				}
				buf.WriteString(translated)
				buf.WriteString("\n")
			case *ast.ThematicBreak:
				log.Debug().Str("Text", string(n.Text(src))).Msg("ast.ThematicBreak")
				buf.WriteString("---")
				buf.WriteString("\n")
			default:
				log.Debug().Str("Type", node.Kind().String()).Str("Text", string(node.Text(src))).Msg("ast.Node [ignored]")
			}
		}
		return ast.WalkContinue, nil
	})

	return buf.String(), nil
}
