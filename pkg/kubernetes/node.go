package kubernetes

import (
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetNodes Return Node List Items
func GetNodes(clientSet *kubernetes.Clientset, labels string) ([]corev1.Node,error) {

	opts := metav1.ListOptions{}
	if labels != "" {
		opts.LabelSelector = labels
	}

	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), opts)
	if err != nil {
		return nil,err
	}

	return nodeList.Items, nil
}
