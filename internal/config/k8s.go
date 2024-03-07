package config

import (
	"flag"
	"log"
	"os"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig() (v1.DeploymentInterface, v1.DaemonSetInterface, v1.ReplicaSetInterface, error) {
	kubeconfig := os.Args[1]
	flag.Parse()
	client, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Println("Error cannot read build config file")
	}

	gotClient, err := kubernetes.NewForConfig(client)
	if err != nil {
		log.Println("Cannot connect to k8s")
	}

	deployment := gotClient.AppsV1().Deployments(corev1.NamespaceDefault)
	daemonset := gotClient.AppsV1().DaemonSets(corev1.NamespaceDefault)
	replica := gotClient.AppsV1().ReplicaSets(corev1.NamespaceDefault)

	return deployment, daemonset, replica, nil

}
