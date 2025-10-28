package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin is running!",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/check", CheckHandler) // call handler in handler.go
	}

}
