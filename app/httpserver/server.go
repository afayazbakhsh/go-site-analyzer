package httpserver

import (
	"github.com/gin-gonic/gin"
)

func Run(addr string) error {
	r := gin.Default()
	RegisterRoutes(r)
	return r.Run(addr)
}
