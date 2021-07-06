package kubernetes

import (
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetServices Return SVC List Items
func GetServices(clientSet *kubernetes.Clientset,listOption metav1.ListOptions,namespace string) ([]corev1.Service,error) {

	svcList, err := clientSet.CoreV1().Services(namespace).List(context.Background(),listOption)
	if err != nil {
		return nil,err
	}
	return svcList.Items, nil
}
