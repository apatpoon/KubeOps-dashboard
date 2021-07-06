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

	api.GET("nodes", controllers.GetNodesListHandler)
	api.GET("deployments", controllers.GetDeploymentListHandler)
	api.GET("services", controllers.GetServiceListHandler)
	api.GET("statefulsets", controllers.GetStatefulSetListHandler)
}
