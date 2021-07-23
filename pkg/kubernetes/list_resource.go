package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func (dp *Deployment) List(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) (interface{}, error) {
	dpList, err := clientSet.AppsV1().Deployments(namespace).List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}
	return dpList.Items, nil
}

func (sts *StatefulSet) List(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) (interface{}, error) {
	stsList, err := clientSet.AppsV1().StatefulSets(namespace).List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}
	return stsList.Items, nil
}

func (svc *Service) List(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) (interface{}, error) {
	svcList, err := clientSet.CoreV1().Services(namespace).List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}
	return svcList.Items, nil
}

func (nd *Node) List(clientSet *kubernetes.Clientset, listOption metav1.ListOptions, namespace string) (interface{}, error) {
	_ = namespace
	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), listOption)
	if err != nil {
		return nil, err
	}

	return nodeList.Items, nil
}
