package controllers

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/pkg/kubernetes"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func GetServiceListHandler(c *gin.Context) {

	listOption := metav1.ListOptions{}
	namespace := c.Request.URL.Query().Get("namespace")
	services, err := kubernetes.GetServices(global.KubernetesClient(), listOption, namespace)
	if err != nil {
		klog.V(2).ErrorS(err, "Getting deployments list failed", "controller", "GetDeploymentHandler")
		WriteErrResp(c, err.Error())
	}
	c.JSON(http.StatusOK, services)
}
