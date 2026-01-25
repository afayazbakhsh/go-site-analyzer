package controllers

import (
	"gocrawler/app/crawler"
	"gocrawler/app/httpserver/requests"
	"net/http"

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
