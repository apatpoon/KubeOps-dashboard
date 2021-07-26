package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Resource interface {
	List(clientSet *kubernetes.Clientset, listOptions metav1.ListOptions, namespace string) (interface{}, error)
	Update(clientSet *kubernetes.Clientset, oldResourceObj interface{}, updateOptions metav1.UpdateOptions) (interface{}, error)
	Delete(clientSet *kubernetes.Clientset, namespace string, name string, deleteOpts metav1.DeleteOptions) error
}

type Deployment struct {
}
type StatefulSet struct {
}
type Service struct {
}
type Node struct {
}
