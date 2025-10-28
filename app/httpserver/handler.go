package httpserver

import (
	"gocrawler/app/crawler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler(c *gin.Context) {
	result := crawler.Run()

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
