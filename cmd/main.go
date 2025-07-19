package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practical/internal/api"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("index.html")
	})
	r.POST("/calculate", api.Processor)
	r.Run()
}
