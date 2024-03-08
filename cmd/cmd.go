package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/rulanugrh/aitne/internal/config"
	"github.com/rulanugrh/aitne/internal/service/apps"
	"github.com/rulanugrh/aitne/internal/service/core"
)

var (
	kubeconfig = flag.String("kubeconfig", "$HOME/.kube/config/", "flag for k8s config file")
	opt        = flag.String("opt", "get", "to use operator ex. get / delete / catch")
	types      = flag.String("types", "deployment", "to use type data ex. deployment / daemonset")
	name       = flag.String("name", "demo-deployment", "this flag use for catch opt ")
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

	pod core.Pods
}

func main() {
	client, err := config.GetConfig(kubeconfig)
	if err != nil {
		log.Printf("Error %s", err.Error())
	}

	newDeployment := apps.NewDeployment(client)
	newDaemonSet := apps.NewDaemonSet(client)
	newReplica := apps.NewReplicaSet(client)
	newPod := core.NewPod(client)

	c := CLI{
		deployment: newDeployment,
		daemon:     newDaemonSet,
		replica:    newReplica,
		pod:        newPod,
	}

	if opt == &get {
		c.get_opt(*types)
	} else if opt == &catch {
		c.catch_opt(*types, *name)
	} else if opt == &deleted {
		c.delete_opt(*types, *name)
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
