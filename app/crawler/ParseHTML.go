package crawler

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedHtml struct {
	Title           string   `json:"title"`
	MetaDescription string   `json:"meta_description"`
	WordCount       int      `json:"word_count"`
	LinksCount      int      `json:"links_count"`
	InternalLinks   []string `json:"internal_links"`
	ExternalLinks   []string `json:"external_links"`
}

func ParseHTML(body []byte, baseURL string) (*ParsedHtml, error) {

	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	desc, _ := doc.Find("meta[name='description']").Attr("content")

	text := doc.Find("body").Text()
	words := strings.Fields(text)
	wordCount := len(words)

	internalLinks, externalLinks := extractLinks(doc, baseURL)

	linksCount := len(internalLinks) + len(externalLinks)

	return &ParsedHtml{
		Title:           title,
		MetaDescription: desc,
		WordCount:       wordCount,
		LinksCount:      linksCount,
		InternalLinks:   internalLinks,
		ExternalLinks:   externalLinks,
	}, nil
}

func extractLinks(doc *goquery.Document, baseURL string) ([]string, []string) {

	internalLinks := make([]string, 0)
	externalLinks := make([]string, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		href, exists := s.Attr("href")

		if exists {
			if strings.Contains(href, baseURL) {
				internalLinks = append(internalLinks, href)
			} else if len(href) > 10 {
				externalLinks = append(externalLinks, href)
			}
		}
	})

	return internalLinks, externalLinks
}
