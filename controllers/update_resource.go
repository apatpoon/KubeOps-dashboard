package controllers

import (
	"KubeOps-dashboard/global"
	"KubeOps-dashboard/pkg/kubernetes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UpdateResourceHandler(c *gin.Context) {

	var ResourceObj kubernetes.Resource

	switch c.Param("resourceType") {
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

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Close Req Body Error: %s\n", err.Error())
		}
	}(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := ResourceObj.Update(global.KubernetesClient(), &bodyBytes, metav1.UpdateOptions{})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
