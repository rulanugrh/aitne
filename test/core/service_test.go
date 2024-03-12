package core

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

type ServiceTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewServiceTest() *ServiceTest {
	return &ServiceTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (service *ServiceTest) TestCreateService() {
	req := model.Service{
		ServiceType: "sample-service",
		Kind:        "Service",
		APIVersion:  "v1",
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName:  "sample-service",
			Annotations: map[string]string{
				"type": "sample-service",
			},
			Labels: map[string]string{
				"type": "sample-service",
			},
		},
		Replica: 2,
		Resource: model.ResourceService{
			ExternalIPs: []string{
				"192.168.100.22",
			},
			ClusterIP:      "172.16.0.1",
			LoadBalancerIP: "192.168.100.1",
			ExternalName:   "sample-service",
			Ports: model.Port{
				Name:     "http",
				Protocol: "tcp",
				NodePort: 8000,
				Port:     80,
			},
			Selector: map[string]string{
				"type": "sample-service",
			},
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := service.client.Post("http://localhost:3000/api/service/create/", byt, service.res)
	if err != nil {
		log.Fatal(err)
	}

	service.Equal(http.StatusOK, resp.StatusCode)
	service.Equal("success create success", res.Message)
}

func (service *ServiceTest) TestGetAllService() {
	res, resp, err := service.client.Get("http://localhost:3000/api/service/getAll/", service.res)
	if err != nil {
		log.Fatal(err)
	}

	service.Equal(http.StatusOK, resp.StatusCode)
	service.Equal("success get all data", res.Message)
}

func (service *ServiceTest) TestGetOneService() {
	res, resp, err := service.client.Get("http://localhost:3000/api/service/get/sample-service", service.res)
	if err != nil {
		log.Fatal(err)
	}

	service.Equal(http.StatusOK, resp.StatusCode)
	service.Equal("success get service by this name", res.Message)
}

func (service *ServiceTest) TestDeleteService() {
	resp, err := service.client.Delete("http://localhost:3000/api/service/delete/sample-service", service.res)
	if err != nil {
		log.Fatal(err)
	}

	service.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestService(t *testing.T) {
	suite.Run(t, NewServiceTest())
}
