package config

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(kubepath *string) (*kubernetes.Clientset, error) {
	client, err := clientcmd.BuildConfigFromFlags("", *kubepath)
	if err != nil {
		log.Println("Error cannot read build config file")
	}

	gotClient, err := kubernetes.NewForConfig(client)
	if err != nil {
		log.Println("Cannot connect to k8s")
	}

	return gotClient, nil

}
