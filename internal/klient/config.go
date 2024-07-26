package klient

import (
	"os"

	"github.com/go-logr/logr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Klient struct {
	logger logr.Logger
	client *kubernetes.Clientset
}

func New(logger logr.Logger) (*Klient, error) {
	klient := &Klient{}
	klient.logger = logger

	kubeConfig, err := getKubeConfig()
	if err != nil {
		return klient, err
	}
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return klient, err
	}
	klient.client = clientset

	return klient, nil
}

func getKubeConfig() (*rest.Config, error) {
	configPath := os.Getenv("KUBECONFIG")

	if configPath != "" {
		co := clientcmd.ConfigOverrides{}
		config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&co).ClientConfig()
		if err != nil {
			return nil, err
		}

		return config, nil
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		return config, nil
	}
}
