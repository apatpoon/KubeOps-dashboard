package kubernetes

import (
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetDeployments Return Deployment List Items
func GetDeployments(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) ([]appsv1.Deployment, error) {
	deployList, err := clientSet.AppsV1().Deployments(namespace).List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}
	return deployList.Items, nil
}
