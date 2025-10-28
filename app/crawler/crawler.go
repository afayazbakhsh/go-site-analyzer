package crawler

import (
	"time"
)

func Crawl(url string) map[string]any {

	result := Fetch(url, 10*time.Second)

	if result.Err != nil {
		return map[string]any{
			"url":   url,
			"error": result.Err.Error(),
		}
	}

	return map[string]any{
		"url":        url,
		"status":     result.StatusCode,
		"load_time":  result.Duration.Seconds(),
		"body_bytes": len(result.Body),
	}
}
