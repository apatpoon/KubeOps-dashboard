package main

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/routers"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main() {

	// Middleware Plugin
	if err := global.Init(); err != nil {
		klog.Error("Init Middleware Plugin Error", err)
		return
	}

	// New A Server
	serv := gin.Default()
	// Init Apis
	routers.InitApi(serv)

	if err := serv.Run(":8088"); err != nil {
		klog.Error("Staring Server Error", err)
	}

}
