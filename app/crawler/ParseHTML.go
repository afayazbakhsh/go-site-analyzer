package crawler

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(body []byte, baseURL string) map[string]any {

	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return map[string]any{"error": err.Error()}
	}

	title := doc.Find("title").Text()
	desc, _ := doc.Find("meta[name='description']").Attr("content")

	text := doc.Text()
	words := strings.Fields(text)
	wordCount := len(words)

	return map[string]any{
		"title":       title,
		"description": desc,
		"word_count":  wordCount,
	}
}
