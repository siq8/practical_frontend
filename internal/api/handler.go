package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"practical/internal/parser"
)

type Request struct {
	Expression string `json:"expression" binding:"required"`
}

func Processor(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := parser.EvalExpr(req.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"expression": req.Expression,
		"result":     result,
	})
} 