package k8s

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client is the k8s-gw k8s client
type Client struct {
	Clientset *kubernetes.Clientset
}

// NewClient initiates a new k8s clientset
func NewClient(kubeconfigPath string, inCluster bool) (*Client, error) {
	var config *rest.Config
	var err error

	if inCluster {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, errors.Wrap(err, "Coud not build in-cluster config")
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, errors.Wrap(err, "Coud not build config from kubeconfigPath")
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create clientset from kubeconfig")
	}
	c := &Client{
		Clientset: client,
	}

	return c, nil
}

// NewClientOrDie returns the clientset or panics
func NewClientOrDie(kubeconfigPath string, inCluster bool) *Client {
	cs, err := NewClient(kubeconfigPath, inCluster)
	if err != nil {
		panic(err.Error())
	}
	return cs
}
