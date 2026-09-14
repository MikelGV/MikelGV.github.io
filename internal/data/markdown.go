package data

import (
	"bytes"
	"html"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
)

// RenderMarkdown converts a Markdown string to a trusted-HTML templ component.
// Content comes from our own blogs.json, so raw HTML output is safe here.
func RenderMarkdown(md string) templ.Component {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		return templ.Raw(html.EscapeString(md))
	}
	return templ.Raw(buf.String())
}
