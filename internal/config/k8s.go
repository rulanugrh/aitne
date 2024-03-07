package config

import (
	"flag"
	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(kubepath *string) (v1.DeploymentInterface, v1.DaemonSetInterface, v1.ReplicaSetInterface, error) {
	flag.Parse()
	client, err := clientcmd.BuildConfigFromFlags("", *kubepath)
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
