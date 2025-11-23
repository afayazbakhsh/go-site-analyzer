package crawler

import (
	"context"
	"io"
	"net"
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
