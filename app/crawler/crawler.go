package crawler

func Crawl(url string) map[string]any {

	result := Fetch(url)

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
