package httpserver

import (
	"github.com/gin-gonic/gin"
	"gocrawler/app/crawler"
	"net/http"
)

func CheckHandler(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing 'url' parameter"})
		return
	}

	result := crawler.Crawl(url)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
