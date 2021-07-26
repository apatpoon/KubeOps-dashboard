package controllers

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/pkg/kubernetes"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListResourceHandler(c *gin.Context) {

	var ResourceObj kubernetes.Resource

	switch c.Param(resourceType) {
	case Deployment:
		ResourceObj = &kubernetes.Deployment{}
	case StatefulSet:
		ResourceObj = &kubernetes.StatefulSet{}
	case Service:
		ResourceObj = &kubernetes.Service{}
	case Node:
		ResourceObj = &kubernetes.Node{}
	default:
		c.JSON(http.StatusBadRequest, fmt.Sprintf("The resource type %s you are searching is not supported", c.Param("resourceType")))
		return
	}

	var namespace string

	if namespace = c.Param("namespace"); len(namespace) <= 0 {
		namespace = corev1.NamespaceDefault
	}

	itemSlice, err := ResourceObj.List(global.KubernetesClient(), metav1.ListOptions{}, namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Getting resource error: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, itemSlice)
}
