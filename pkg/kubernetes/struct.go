package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Resource interface {
	List(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) (interface{}, error)
}

type Deployment struct {
}
type StatefulSet struct {
}
type Service struct {
}
type Node struct {
}
