package crawler

import (
	"bytes"
	"fmt"
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

	text := doc.Find("body").Text()
	words := strings.Fields(text)
	wordCount := len(words)

	links := extractLinks(doc)

	return map[string]any{
		"title":            title,
		"meta_description": desc,
		"word_count":       wordCount,
		"links_count":      len(links),
	}
}

func extractLinks(doc *goquery.Document) []string {
	links := make([]string, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println("Extracting links...", href)

			links = append(links, href)
		}
	})

	return links
}
