package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"bytes"

	"github.com/117503445/markdown-translate/test/examples"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestMarkdown(t *testing.T) {
	source := []byte("# Hello\n\nThis is a markdown file.")

	r, err := translateMarkdown(source)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Log(r)

    r , _ = translateMarkdown([]byte(examples.Mempool))

    os.WriteFile("mempool.md", []byte(r), 0644)


    t.Log(examples.Mempool)
}

func translateText(lang, text string) (string, error) {
	return fmt.Sprintf("[翻译]%s[结束]", text), nil
}

func translateMarkdown(md []byte) (string, error) {
	var buf bytes.Buffer

	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	doc := markdown.Parser().Parse(text.NewReader(md))

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
        if entering {
            switch n := node.(type) {
            case *ast.Heading:
                level := n.Level
                buf.WriteString(strings.Repeat("#", level) + " ")
                buf.WriteString(string(n.Text(md)))
                buf.WriteString(" ")
                translated, err := translateText("zh", string(n.Text(md)))
                if err != nil {
                    return ast.WalkStop, err
                }
                buf.WriteString(translated)
                buf.WriteString("\n")
                return ast.WalkSkipChildren, nil
            case *ast.Paragraph:
                buf.WriteString("\n")
            case *ast.Text:
                buf.WriteString(string(n.Text(md)))
                buf.WriteString("\n\n")
                translated, err := translateText("zh", string(n.Text(md)))
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
