package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rulanugrh/aitne/internal/config"
	"github.com/rulanugrh/aitne/internal/service/apps"
	"github.com/rulanugrh/aitne/internal/service/core"
)

var (
	kubeconfig = os.Getenv("KUBECONFIG_PATH")
	opt        = os.Args[0]
	types      = os.Args[1]
	name       = flag.String("name", "sample-deployment", "this command for operate by name")
)

var (
	get     = "get"
	catch   = "catch"
	deleted = "delete"
)

type CLI struct {
	deployment apps.Deployment
	daemon     apps.DaemonSet
	replica    apps.ReplicaSet

	pod       core.Pods
	sevice    core.Service
	namespace core.Namespace
	configmap core.ConfigMaps
	replicac  core.ReplicationController
	node      core.Node
	secret    core.Secret
}

func help() {
	helpContent := [][]string {
		{ "help", "show help message" },
		{ "get" , "type of get data " },
		{ "catch", "type of to catch data by name"  },
		{ "delete", "type of to delete data by name" },
		{ "[type]", "type operator like deployment / daemon / etc" },
		{ "-name", "name of services from type" },
	}

	example := "\nexample: go run cmd/cmd.go get deployment"

	max := len(helpContent[0][0])
	for _, part := range helpContent {
		length := len(part[0])
		if length > max {
			max = length
		}
	}

	var builder strings.Builder
	const space = 4

	for _, part := range helpContent {
		builder.WriteString(part[0])
		spacer := (max - len(part[0])) + space
		for spacer > 0 {
			builder.WriteByte(' ')
			spacer--
		}

		builder.WriteString(part[1])
		builder.WriteByte('\n')
	}

	println(builder.String()[:builder.Len()-1])
	println(example)
}

func main() {
	client, err := config.GetConfig(&kubeconfig)
	if err != nil {
		log.Printf("Error %s", err.Error())
	}

	c := CLI{
		deployment: apps.NewDeployment(client),
		daemon:     apps.NewDaemonSet(client),
		replica:    apps.NewReplicaSet(client),
		pod:        core.NewPod(client),
		sevice:     core.NewServiceKurbenetes(client),
		namespace:  core.NewNamespace(client),
		replicac:   core.NewReplicationController(client),
		configmap:  core.NewConfigMap(client),
		secret:     core.NewSecretKurbenetes(client),
		node:       core.NewNodeConfig(client),
	}

	if opt == "help" {
		help()
	} else if opt == get {
		c.get_opt(types)
	} else if opt == catch {
		c.catch_opt(types, *name)
	} else if opt == deleted {
		c.delete_opt(types, *name)
	} else {
		log.Println("your opt is invalid, use go run cmd.go -h to see details")
	}
}

func (c *CLI) get_opt(types string) {
	if types == "deployment" {
		convert_response_to_json(c.deployment.List())
	} else if types == "daemon" {
		convert_response_to_json(c.daemon.List())
	} else if types == "replica" {
		convert_response_to_json(c.replica.List())
	} else if types == "pod" {
		convert_response_to_json(c.pod.List())
	} else if types == "service" {
		convert_response_to_json(c.sevice.List())
	} else if types == "namespace" {
		convert_response_to_json(c.namespace.List())
	} else if types == "replicac" {
		convert_response_to_json(c.replicac.List())
	} else if types == "configmap" {
		convert_response_to_json(c.configmap.List())
	} else if types == "node" {
		convert_response_to_json(c.node.List())
	} else if types == "secret" {
		convert_response_to_json(c.secret.List())
	} else {
		log.Println("sorry invalid type data")
	}
}

func (c *CLI) catch_opt(types string, name string) {
	if types == "deployment" {
		convert_response_to_json(c.deployment.GetByName(name))
	} else if types == "daemon" {
		convert_response_to_json(c.daemon.GetByName(name))
	} else if types == "pod" {
		convert_response_to_json(c.pod.GetByName(name))
	} else if types == "replica" {
		convert_response_to_json(c.replica.GetByName(name))
	} else if types == "service" {
		convert_response_to_json(c.sevice.GetByName(name))
	} else if types == "namespace" {
		convert_response_to_json(c.namespace.GetByName(name))
	} else if types == "replicac" {
		convert_response_to_json(c.replicac.GetByName(name))
	} else if types == "configmap" {
		convert_response_to_json(c.configmap.GetByName(name))
	} else if types == "node" {
		convert_response_to_json(c.node.GetByName(name))
	} else if types == "secret" {
		convert_response_to_json(c.secret.GetByName(name))
	} else {
		log.Println("sorry invalid type data")
	}
}

func (c *CLI) delete_opt(types string, name string) {
	if types == "deployment" {
		printout_response(c.deployment.Delete(name))
	} else if types == "daemon" {
		printout_response(c.daemon.Delete(name))
	} else if types == "replica" {
		printout_response(c.replica.Delete(name))
	} else if types == "pod" {
		printout_response(c.pod.Delete(name))
	} else if types == "service" {
		printout_response(c.sevice.Delete(name))
	} else if types == "namespace" {
		printout_response(c.namespace.Delete(name))
	} else if types == "replicac" {
		printout_response(c.replicac.Delete(name))
	} else if types == "configmap" {
		printout_response(c.configmap.Delete(name))
	} else if types == "node" {
		printout_response(c.node.Delete(name))
	} else if types == "secert" {
		printout_response(c.secret.Delete(name))
	} else {
		log.Println("invalid type data")
	}
}

func convert_response_to_json(data any, err error) {
	if err != nil {
		log.Printf("something error %s", err.Error())
	}

	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("errror marshal json %s", err.Error())
	}

	fmt.Println(string(response))

}

func printout_response(err error) {
	if err != nil {
		log.Printf("something error %s", err.Error())
	}

	log.Println("success operator")
}
