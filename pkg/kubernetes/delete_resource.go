package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func (dp *Deployment) Delete(clientSet *kubernetes.Clientset, namespace string, resourceName string, deleteOpts metav1.DeleteOptions) error {

	err := clientSet.AppsV1().Deployments(namespace).Delete(context.Background(), resourceName, deleteOpts)
	if err != nil {
		return err
	}
	return nil
}

func (sts *StatefulSet) Delete(clientSet *kubernetes.Clientset, namespace string, name string, deleteOpts metav1.DeleteOptions) error {
	err := clientSet.AppsV1().StatefulSets(namespace).Delete(context.Background(), name, deleteOpts)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Delete(clientSet *kubernetes.Clientset, namespace string, name string, deleteOpts metav1.DeleteOptions) error {
	err := clientSet.CoreV1().Services(namespace).Delete(context.Background(), name, deleteOpts)
	if err != nil {
		return err
	}
	return nil
}

func (nd *Node) Delete(clientSet *kubernetes.Clientset, namespace string, name string, deleteOpts metav1.DeleteOptions) error {
	err := clientSet.CoreV1().Nodes().Delete(context.Background(), name, deleteOpts)
	if err != nil {
		return err
	}
	return nil
}
