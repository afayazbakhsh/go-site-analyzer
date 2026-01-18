package httpserver

import (
	"github.com/gin-gonic/gin"
)

func Run(addr string) error {

	// r := gin.Default() پیش فرض به همراه logger & recovery

	r := gin.New()        // بدون میدلورهای پیش‌فرض
	r.Use(gin.Logger())   // اضافه کردن Logger به‌صورت دلخواه
	r.Use(gin.Recovery()) // اضافه کردن Recovery

	r.Use(adminAuthMiddlewares())

	RegisterRoutes(r) // جنریت کردن روت های برنامه

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "route not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "method not allowed"})
	})

	return r.Run(addr) // راه اندازی روی پورت مورد نظر
}
