package kubernetes

import (
	"context"
	"encoding/json"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (dp *Deployment) Update(clientSet *kubernetes.Clientset, bodyBytes interface{}, updateOptions metav1.UpdateOptions) (interface{}, error) {

	var newDeploy *appsv1.Deployment

	err := json.Unmarshal(*bodyBytes.(*[]byte), &newDeploy)

	if err != nil {
		return nil, err
	}

	oldDeploy, err := clientSet.AppsV1().Deployments(newDeploy.Namespace).Get(context.Background(), newDeploy.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	oldDeploy.Spec = newDeploy.Spec

	newDeploy, err = clientSet.AppsV1().Deployments(newDeploy.Namespace).Update(context.Background(), oldDeploy, updateOptions)

	if err != nil {
		return nil, err
	}

	return newDeploy, nil
}

func (sts *StatefulSet) Update(clientSet *kubernetes.Clientset, oldResourceObj interface{}, updateOptions metav1.UpdateOptions) (interface{}, error) {

	return nil, nil
}

func (svc *Service) Update(clientSet *kubernetes.Clientset, oldResourceObj interface{}, updateOptions metav1.UpdateOptions) (interface{}, error) {

	return nil, nil
}

func (nd *Node) Update(clientSet *kubernetes.Clientset, oldResourceObj interface{}, updateOptions metav1.UpdateOptions) (interface{}, error) {

	return nil, nil
}
