package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/rulanugrh/aitne/internal/config"
	"github.com/rulanugrh/aitne/internal/service"
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
	deployment service.Deployment
	daemon     service.DaemonSet
	replica    service.ReplicaSet
}

func main() {
	deployment, daemon, replica, _, err := config.GetConfig(kubeconfig)
	if err != nil {
		log.Printf("Error %s", err.Error())
	}

	newDeployment := service.NewDeployment(deployment)
	newDaemonSet := service.NewDaemonSet(daemon)
	newReplica := service.NewReplicaSet(replica)

	c := CLI{
		deployment: newDeployment,
		daemon:     newDaemonSet,
		replica:    newReplica,
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
	} else {
		log.Println("sorry invalid type data")
	}
}

func (c *CLI) catch_opt(types string, name string) {
	if types == "deployment" {
		convert_response_to_json(c.deployment.GetByName(name))
	} else if types == "daemon" {
		convert_response_to_json(c.daemon.GetByName(name))
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
