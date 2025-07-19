package main

import (
	"github.com/gin-gonic/gin"
	"practical/internal/api"
)

func main() {
	r := gin.Default()
	r.POST("/calculate", api.Processor)
	r.Run()
}
