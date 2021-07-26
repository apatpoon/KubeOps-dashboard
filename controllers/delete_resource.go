package controllers

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/pkg/kubernetes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gin-gonic/gin"
)

func DeleteResourceHandler(c *gin.Context) {
	var ResourceObj kubernetes.Resource

	switch c.Param(resourceType) {
	case Deployment:
		ResourceObj = &kubernetes.Deployment{}
	case StatefulSet:
		ResourceObj = &kubernetes.StatefulSet{}
	case Service:
		ResourceObj = &kubernetes.Service{}
	default:
		c.JSON(http.StatusBadRequest, fmt.Sprintf("The resource type %s you are searching is not supported", c.Param("resourceType")))
		return
	}

	reqParam := make(map[string]string)

	var namespace string

	bytes, _ := ioutil.ReadAll(c.Request.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Close Req Body Error: %s\n", err.Error())
		}
	}(c.Request.Body)

	err := json.Unmarshal(bytes, &reqParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Unmarshal Error: %s", err.Error()))
		return
	}

	resourceName := reqParam["name"]
	if len(resourceName) <= 0 {
		c.JSON(http.StatusBadRequest, "Resource name required")
		return
	}
	if namespace = reqParam["namespace"]; len(namespace) <= 0 {
		namespace = corev1.NamespaceDefault
	}

	err = ResourceObj.Delete(global.KubernetesClient(), namespace, resourceName, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Delete resource failed: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, "OK")
}
