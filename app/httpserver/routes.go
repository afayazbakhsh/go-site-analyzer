package httpserver

import (
	"gocrawler/app/httpserver/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin is running!",
		})
	})

	api := route.Group("/api/v1/")
	{
		api.GET("check", controllers.CheckHandler) // call handler in handler.go

		api.Group("page-data/")
		{
			api.POST("create", controllers.CreatePageData)
		}
	}
}
