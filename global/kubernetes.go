package global

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
)

var client *kubernetes.Clientset

func InitKubernetesClient() error {
	var err error
	var config 	*rest.Config

	// InCluster(pod) Or KubeConfig(kubectl)
	kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	if config, err = rest.InClusterConfig(); err != nil {
		klog.V(2).ErrorS(err,"Not in Cluster")
		if config, err = clientcmd.BuildConfigFromFlags("",kubeConfig); err != nil {
			return err
		}
	}

	// Successfully Get A Kubernetes Cluster Config
	if client, err = kubernetes.NewForConfig(config); err != nil {
		return err
	}
	return nil
}

func KubernetesClient() *kubernetes.Clientset {
	return client
}