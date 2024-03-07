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
	opt   = flag.String("opt", "get", "to use operator ex. get / delete / catch")
	types = flag.String("types", "deployment", "to use type data ex. deployment / daemonset")
	name  = flag.String("name", "demo-deployment", "this flag use for catch opt ")
)

var (
	get     = "get"
	catch   = "catch"
	deleted = "delete"
)

type CLI struct {
	deployment service.Deployment
	daemon     service.DaemonSet
}

func main() {
	deployment, daemon, _, err := config.GetConfig()
	if err != nil {
		log.Printf("Error %s", err.Error())
	}

	newDeployment := service.NewDeployment(deployment)
	newDaemonSet := service.NewDaemonSet(daemon)
	c := CLI{
		deployment: newDeployment,
		daemon:     newDaemonSet,
	}

	if opt == &get {
		c.get_opt(*types)
	} else if opt == &catch {
		c.catch_opt(*types, *name)
	} else {
		log.Println("your opt is invalid, use go run cmd.go -h to see details")
	}
}

func (c *CLI) get_opt(types string) {
	if types == "deployment" {
		convert_response_to_json(c.deployment.List())
	} else if types == "daemon" {
		convert_response_to_json(c.daemon.List())
	} else {
		log.Println("sorry invalid type data")
	}
}

func (c *CLI) catch_opt(types string, name string) {
	if types == "deployment" {
		convert_response_to_json(c.deployment.GetByName(name))
	} else if types == "daemon" {
		convert_response_to_json(c.daemon.GetByName(name))
	} else {
		log.Println("sorry invalid type data")
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
