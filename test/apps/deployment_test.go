package apps

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/util/constant"
	helper "github.com/rulanugrh/aitne/test"
	"github.com/stretchr/testify/suite"
)

type DeploymentTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewDeploymentTest() *DeploymentTest {
	return &DeploymentTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (deployment *DeploymentTest) TestCreateDeployment() {
	req := model.CreateDeployment{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
		MatchLabels: map[string]string{
			"type": "demo-deployment",
		},
		Container: model.Container{
			Name:         "web-app",
			Image:        "nginx:latest",
			PortExposed:  8080,
			Protocol:     "http",
			NameProtocol: "tcp",
		},
		Labels: map[string]string{
			"app": "demo",
		},
		Annotations: map[string]string{
			"app": "demo",
		},
		Name: "demo-deployment",
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := deployment.client.Post("http://localhost:3000/api/deployment/create/", byt, deployment.res)
	if err != nil {
		log.Fatal(err)
	}

	deployment.Equal(http.StatusOK, resp.StatusCode)
	deployment.Equal("success create deployment", res.Message)
}

func (deployment *DeploymentTest) TestGetAllDeployment() {
	res, resp, err := deployment.client.Get("http://localhost:3000/api/deployment/getAll/", deployment.res)
	if err != nil {
		log.Fatal(err)
	}

	deployment.Equal(http.StatusOK, resp.StatusCode)
	deployment.Equal("success get all data", res.Message)
}

func (deployment *DeploymentTest) TestGetOneDeployment() {
	res, resp, err := deployment.client.Get("http://localhost:3000/api/deployment/get/demo-deployment", deployment.res)
	if err != nil {
		log.Fatal(err)
	}

	deployment.Equal(http.StatusOK, resp.StatusCode)
	deployment.Equal("success get deployment by this name", res.Message)
}

func (deployment *DeploymentTest) TestDeleteDeployment() {
	resp, err := deployment.client.Delete("http://localhost:3000/api/deployment/delete/demo-deployment", deployment.res)
	if err != nil {
		log.Fatal(err)
	}

	deployment.Equal(http.StatusOK, resp.StatusCode)
}

func TestDeployment(t *testing.T) {
	suite.Run(t, NewDeploymentTest())
}
