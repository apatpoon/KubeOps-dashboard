package controllers

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/pkg/kubernetes"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

func GetNodesListHandler (c *gin.Context) {
	nodes, err := kubernetes.GetNodes(global.KubernetesClient(),"")
	if err != nil {
		klog.V(2).ErrorS(err,"Getting node list failed", "controller", "GetNodeListHandler")
		WriteErrResp(c,err.Error())
	}
	c.JSON(http.StatusOK, nodes)
}