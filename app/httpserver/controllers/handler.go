package controllers

import (
	"fmt"
	"gocrawler/app/crawler"
	"gocrawler/app/db"
	"gocrawler/app/httpserver/requests"
	"gocrawler/app/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func CheckHandler(c *gin.Context) {

	var requests requests.UrlCheckRequest

	if err := c.ShouldBind(&requests); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	readResult, err := crawler.Read(*requests.Url)

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to read Data", "error": err})
		return
	}

	writeResult, err := crawler.Write(readResult)

	if err != nil {
		c.JSON(500, gin.H{"message": "failed to write crawl result", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"read":  readResult,
		"write": writeResult,
	})
}


func CheckMainPagesData() {
	fmt.Println("Start scraping...")

	URLs := []string{
		"https://dojinja.com/",
		"https://eghtesadkhabar.com/",
		"http://qamarnews.com/",
	}

	// Channels
	readPages := make(chan *crawler.ReadPage, 5) // buffered for backpressure
	errorsChan := make(chan string, 5)

	var wg sync.WaitGroup       // for scraping
	var dbWg sync.WaitGroup     // for DB workers

	// --- DB workers ---
	numDBWorkers := 2 // adjust as needed
	dbWg.Add(numDBWorkers)
	for i := 0; i < numDBWorkers; i++ {
		
		go func(id int) {
			defer dbWg.Done()
			for page := range readPages {
				// Call Write method to save/update in DB
				if _, err := crawler.Write(page); err != nil {
					fmt.Printf("[DB Worker %d] Failed to save page %s: %v\n", id, page.URL, err)
				} else {
					fmt.Printf("[DB Worker %d] Saved page: %s\n", id, page.URL)
				}
			}
		}(i)
	}

	// --- Scraper goroutines ---
	for _, url := range URLs {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			page, err := crawler.Read(u)
			if err != nil {
				errorsChan <- fmt.Sprintf("Error reading %s: %v", u, err)
				return
			}

			readPages <- page // send to DB worker
			fmt.Println("Scraped:", u)
		}(url)
	}

	// --- Close channels when scraping finishes ---
	go func() {
		wg.Wait()          // wait for all scrapers
		close(readPages)   // signal DB workers no more pages
		close(errorsChan)  // signal no more errors
	}()

	// --- Handle errors while scraping ---
	for err := range errorsChan {
		fmt.Println("Error:", err)
	}

	// --- Wait for DB workers to finish ---
	dbWg.Wait()

	fmt.Println("All done!")
}

