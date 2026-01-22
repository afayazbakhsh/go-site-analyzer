package crawler

import (
	"context"
	"gocrawler/app/db"
	"gocrawler/app/models"
	"io"
	"net"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type FetchResult struct {
	Body       []byte
	StatusCode int
	Duration   time.Duration
	Err        error
}

type Links struct {
	Internal []string `json:"internal"`
	External []string `json:"external"`
}

type ReadPage struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WordCount   int    `json:"word_count"`
	LinksCount  int    `json:"links_count"`
	Links       Links  `json:"links"`
	StatusCode  int    `json:"status_code"`
	LoadTime    int64  `json:"load_time"`
}

func Read(url string) (_ *ReadPage, err error) {

	result := Fetch(url) // step 1: request and fetch data

	if result.Err != nil {
		return nil, result.Err
	}

	parsedData, err := ParseHTML(result.Body, url) // step 2: parse the html body

	if err != nil {
		return nil, err
	}

	return &ReadPage{
		URL:         url,
		Title:       parsedData.Title,
		Description: parsedData.MetaDescription,
		WordCount:   parsedData.WordCount,
		LinksCount:  parsedData.LinksCount,
		Links: Links{
			Internal: parsedData.InternalLinks,
			External: parsedData.ExternalLinks,
		},
		StatusCode: result.StatusCode,
		LoadTime:   result.Duration.Milliseconds(),
	}, nil
}

func Write(readData *ReadPage) (*models.PageData, error) {

	var pageData models.PageData

	if err := db.DB.Where("url", readData.URL).First(&pageData).Error; err != nil {

		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		pageData = models.PageData{
			URL:         readData.URL,
			Title:       readData.Title,
			Description: readData.Description,
			WordCount:   readData.WordCount,
			Links: models.PageLinks{
				Internal: len(readData.Links.Internal),
				External: len(readData.Links.External),
			},
			StatusCode: readData.StatusCode,
			LoadTime:   readData.LoadTime,
		}

		if err := db.DB.Create(&pageData).Error; err != nil {
			return nil, err
		}

	} else {

		pageData.Title = readData.Title
		pageData.Description = readData.Description
		pageData.WordCount = readData.WordCount
		pageData.Links = models.PageLinks{
			Internal: len(readData.Links.Internal),
			External: len(readData.Links.External),
		}
		pageData.StatusCode = readData.StatusCode
		pageData.LoadTime = readData.LoadTime
		pageData.UpdatedAt = time.Now()

		if err := db.DB.Where("url", readData.URL).Save(&pageData).Error; err != nil {
			return nil, err
		}
	}

	return &pageData, nil
}

func Fetch(url string) FetchResult {

	start := time.Now()
	// Dial + TLS timeout
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
	}

	// Whole request timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return FetchResult{Err: err}
	}

	resp, err := client.Do(req)
	if err != nil {
		return FetchResult{Err: err}
	}

	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return FetchResult{Err: err}
	}

	return FetchResult{
		Body:       body,
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
		Err:        nil,
	}
}
