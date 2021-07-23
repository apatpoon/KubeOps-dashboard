package routers

import (
	"KubeOps-dashboard/controllers"
	"KubeOps-dashboard/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func InitApi(eng *gin.Engine) {
	// Check Health Api
	eng.GET("/ping", controllers.Ping)

	// Apis Group
	api := eng.Group("/api/v1")

	// Registering Middleware Here
	eng.Use(middlewares.CorsMiddleware)

	api.GET("/:resourceType", controllers.ListResourceHandler)
	api.POST("/:resourceType", controllers.UpdateResourceHandler)
}
