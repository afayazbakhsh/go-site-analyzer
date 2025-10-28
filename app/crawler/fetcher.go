package crawler

import (
	"io"
	"net/http"
	"time"
)

// FetchResult holds the result of a fetch
type FetchResult struct {
	Body       []byte
	StatusCode int
	Duration   time.Duration
	Err        error
}

// Fetch fetches a URL and measures how long it takes
func Fetch(url string, timeout time.Duration) FetchResult {
	start := time.Now()

	client := http.Client{Timeout: timeout}
	response, err := client.Get(url)

	if err != nil {
		return FetchResult{Err: err}
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return FetchResult{Err: err}
	}

	return FetchResult{
		Body:       body,
		StatusCode: response.StatusCode,
		Duration:   time.Since(start),
		Err:        nil,
	}
}
