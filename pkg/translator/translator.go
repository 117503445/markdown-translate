package translator

import (
	"bytes"
	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"strings"
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

func getListItemText(node ast.Node, src []byte) string {
	if node.Kind() != ast.KindListItem {
		panic("node is not a ListItem")
	}
	rawText := ""
	child := node.FirstChild()
	for child != nil {
		rawText = getRawText(child, src)
		child = child.NextSibling()
	}

	return rawText
}

func (t *Translator) Translate(source string) (string, error) {
	var buf bytes.Buffer

	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	src := []byte(source)

	doc := markdown.Parser().Parse(text.NewReader(src))

	// doc.Dump(src, 2)

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {

		s := ""

		if entering {
			// log.Debug().Str("Type", node.Kind().String()).Msg("Node")
			switch n := node.(type) {
			case *ast.Heading:
				level := n.Level
				s += strings.Repeat("#", level) + " "
				s += string(n.Text(src)) + " "
				translated, err :=
					t.provider.Translate(string(n.Text(src)))
				if err != nil {
					return ast.WalkStop, err
				}
				s += translated + "\n\n"
			case *ast.ThematicBreak:
				s += "---\n"
			case *ast.Paragraph:
				raw := getRawText(n, src)
				s += raw + "\n\n"

				translated, err := t.provider.Translate(raw)
				if err != nil {
					return ast.WalkStop, err
				}

				s += translated + "\n\n"
			case *ast.List:
				rawS := ""

				child := n.FirstChild()
				for child != nil {
					switch c := child.(type) {
					case *ast.ListItem:
						raw := getListItemText(c, src)
						rawS += "- " + raw + "\n"
					default:
						log.Warn().Str("Type", c.Kind().String()).Msg("ignore Node in List")
					}
					child = child.NextSibling()
				}

				s += rawS + "\n"

				translated, err := t.provider.Translate(rawS)
				if err != nil {
					return ast.WalkStop, err
				}

				s += translated + "\n"

			case *ast.FencedCodeBlock:
				raw := getRawText(n, src)
				s += "```" + string(n.Language(src)) + "\n" + raw + "```\n"

			case *ast.HTMLBlock:
				raw := getRawText(n, src)
				s += raw + "\n"

			case *ast.Document:
				return ast.WalkContinue, nil
			default:
				log.Warn().Str("Type", node.Kind().String()).Str("Text", string(node.Text(src))).Str("Raw", getRawText(node, src)).
					Msg("ast.Node [ignored]")
			}
		}
		_, err := buf.WriteString(s)
		if err != nil {
			return ast.WalkStop, err
		}
		return ast.WalkSkipChildren, nil
	})

	return buf.String(), nil
}
