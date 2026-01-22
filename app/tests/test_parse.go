package tests

import (
	"fmt"
	"gocrawler/app/crawler"
)

func TestParseHtml() {
	url := "https://eghtesadkhabar.com/%d9%82%db%8c%d9%85%d8%aa-%d8%b7%d9%84%d8%a7-%d9%88-%d8%b3%da%a9%d9%87-%d8%a7%d9%85%d8%b1%d9%88%d8%b2-%db%8c%da%a9%d8%b4%d9%86%d8%a8%d9%87-16-%d8%a2%d8%b0%d8%b1-%d8%a7%d9%81%d8%b2%d8%a7%db%8c%d8%b4/"

	// Fetch the real URL
	fetchResult := crawler.Fetch(url)
	if fetchResult.Err != nil {
		fmt.Printf("Fetch error: %v\n", fetchResult.Err)
		return
	}

	// Parse the HTML
	parsed := crawler.ParseHTML(fetchResult.Body, url)

	// Dump the result
	fmt.Printf("Parsed result: %+v\n", parsed)
}
