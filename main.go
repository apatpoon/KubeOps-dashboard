package main

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/routers"
	"flag"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main() {

	// Init klog
	klog.InitFlags(nil)
	defer klog.Flush()
	if err := flag.Set("logtostderr", "true"); err != nil {
		panic(err)
	}
	if err := flag.Set("alsologtostderr", "true"); err != nil {
		panic(err)
	}
	flag.Parse()

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
