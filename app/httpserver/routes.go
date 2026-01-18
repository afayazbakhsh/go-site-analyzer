package httpserver

import (
	"gocrawler/app/httpserver/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin is running!")
	})

	api := route.Group("/api/v1/")
	{
		api.GET("check", controllers.CheckHandler) // call handler in handler.go

		pagedata := api.Group("page-data/", adminAuthMiddlewares())
		{
			pagedata.GET("index", controllers.Index)
			pagedata.GET(":id/show", controllers.Show)
			pagedata.POST("create", controllers.Create)
			pagedata.PUT("update", controllers.Update)
			pagedata.DELETE("delete", controllers.Delete)
		}
	}
}

func adminAuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		// token := c.GetHeader("Authorization")

		// if token != "secret-token" {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }

		c.Next()
	}
}
