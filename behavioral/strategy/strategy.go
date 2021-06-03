package main

import (
  "fmt"
  "strings"
)

// OutputFormat ...
type OutputFormat int

const (
  // Markdown ...
  Markdown OutputFormat = iota
  // HTML ...
  HTML
)

// ListStrategy ...
type ListStrategy interface {
  Start(builder *strings.Builder)
  End(builder *strings.Builder)
  AddListItem(builder *strings.Builder, item string)
}

// MarkdownListStrategy ...
type MarkdownListStrategy struct {}

// Start ...
func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

// End ...
func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}

// AddListItem ...
func (m *MarkdownListStrategy) AddListItem(
  builder *strings.Builder, item string) {
  builder.WriteString(" * " + item + "\n")
}

// HTMLListStrategy ...
type HTMLListStrategy struct {}

// Start ...
func (h *HTMLListStrategy) Start(builder *strings.Builder) {
  builder.WriteString("<ul>\n")
}

// End ...
func (h *HTMLListStrategy) End(builder *strings.Builder) {
  builder.WriteString("</ul>\n")
}

// AddListItem ...
func (h *HTMLListStrategy) AddListItem(builder *strings.Builder, item string) {
  builder.WriteString("  <li>" + item + "</li>\n")
}

// TextProcessor ...
type TextProcessor struct {
  builder strings.Builder
  listStrategy ListStrategy
}

// NewTextProcessor ...
func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
  return &TextProcessor{strings.Builder{}, listStrategy}
}

// SetOutputFormat ...
func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
  switch fmt {
  case Markdown:
    t.listStrategy = &MarkdownListStrategy{}
  case HTML:
    t.listStrategy = &HTMLListStrategy{}
  }
}

// AppendList ...
func (t *TextProcessor) AppendList(items []string) {
  t.listStrategy.Start(&t.builder)
  for _, item := range items {
    t.listStrategy.AddListItem(&t.builder, item)
  }
  t.listStrategy.End(&t.builder)
}

// Reset ...
func (t *TextProcessor) Reset() {
  t.builder.Reset()
}

// String ...
func (t *TextProcessor) String() string {
  return t.builder.String()
}

func main() {
  tp := NewTextProcessor(&MarkdownListStrategy{})
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)

  tp.Reset()
  tp.SetOutputFormat(HTML)
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)
}