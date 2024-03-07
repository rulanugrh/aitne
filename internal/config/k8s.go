package config

import (
	"flag"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetConfig() (v1.DeploymentInterface, error) {
	var kubeconfig *string
	home := homedir.HomeDir()
	if home != "" {
		kubeconfig = flag.String("config", filepath.Join(home, ".kube", "config"), "file config absolute path")
	} else {
		kubeconfig = flag.String("config", "", "absolute path to file config")
	}

	flag.Parse()
	client, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println("Error cannot read build config file")
	}

	gotClient, err := kubernetes.NewForConfig(client)
	if err != nil {
		log.Println("Cannot connect to k8s")
	}

	deployment := gotClient.AppsV1().Deployments(corev1.NamespaceDefault)
	return deployment, nil

}
