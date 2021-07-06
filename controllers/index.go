package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// WriteErrResp TODO Business Error Code
func WriteErrResp(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"message": msg,
	})
}