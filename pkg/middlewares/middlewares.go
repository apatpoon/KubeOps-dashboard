package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware 处理跨域插件
func CorsMiddleware(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, access-control-allow-origin, access-control-allow-headers")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	if method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
