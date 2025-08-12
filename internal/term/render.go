package term

import (
	"fmt"
	"strings"

	"github.com/rosikui/code-quote/internal/quote"
)

// Renderer handles quote output formatting
type Renderer struct {
	noColor  bool
	markdown bool
	showTags bool
}

// NewRenderer creates a new renderer instance
func NewRenderer(noColor, markdown, showTags bool) *Renderer {
	return &Renderer{
		noColor:  noColor,
		markdown: markdown,
		showTags: showTags,
	}
}

// Render formats and outputs a quote
func (r *Renderer) Render(q *quote.Quote) string {
	if r.markdown {
		return r.renderMarkdown(q)
	}
	return r.renderTerminal(q)
}

// renderTerminal renders quote for terminal output with ANSI colors
func (r *Renderer) renderTerminal(q *quote.Quote) string {
	if r.noColor {
		return r.renderPlain(q)
	}

	// ANSI color codes: bright purple bold
	const (
		colorStart = "\x1b[95;1m"
		colorEnd   = "\x1b[0m"
	)

	text := fmt.Sprintf("%s%s%s", colorStart, q.Text, colorEnd)
	author := fmt.Sprintf("%s%s%s", colorStart, q.Author, colorEnd)

	if r.showTags && len(q.Tags) > 0 {
		tags := strings.Join(q.Tags, ", ")
		return fmt.Sprintf("%s\n— %s [%s]", text, author, tags)
	}

	return fmt.Sprintf("%s\n— %s", text, author)
}

// renderPlain renders quote without colors
func (r *Renderer) renderPlain(q *quote.Quote) string {
	if r.showTags && len(q.Tags) > 0 {
		tags := strings.Join(q.Tags, ", ")
		return fmt.Sprintf("%s\n— %s [%s]", q.Text, q.Author, tags)
	}

	return fmt.Sprintf("%s\n— %s", q.Text, q.Author)
}

// renderMarkdown renders quote in Markdown blockquote format
func (r *Renderer) renderMarkdown(q *quote.Quote) string {
	if r.showTags && len(q.Tags) > 0 {
		tags := strings.Join(q.Tags, ", ")
		return fmt.Sprintf("> %s\n> — %s [%s]", q.Text, q.Author, tags)
	}

	return fmt.Sprintf("> %s\n> — %s", q.Text, q.Author)
}
