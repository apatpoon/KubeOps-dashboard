package kubernetes

import (
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetStatefulSets Return StatefulSet List Items
func GetStatefulSets(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) ([]appsv1.StatefulSet, error) {

	statefulsetList, err := clientSet.AppsV1().StatefulSets(namespace).List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}
	return statefulsetList.Items, nil
}
