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

type ReplicaControllerTest struct {
	suite.Suite
	client helper.SuiteInterface
	res *constant.Response
}

func NewReplicaControllerTest() *ReplicaControllerTest {
	return &ReplicaControllerTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res: &constant.Response{},
	}
}

func (replica *ReplicaControllerTest) TestCreateReplica() {
	req := model.ReplicationController{
		APIVersion: "v1",
		Kind: "ReplicationController",
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName: "sample-replicac",
			Annotations: map[string]string{
				"type": "replication-controller",
			},
			Labels: map[string]string{
				"env": "dev",
			},
		},
		Replica: 2,
		MinReady: 60,
		Selector: map[string]string{
			"type": "replicac",
		},
		Container: model.Container{
			Name: "web-nginx",
			Image: "nginx:latest",
			PortExposed: 3000,
			NameProtocol: "http",
			Protocol: "tcp",
		},
	}

	jsonByte, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byts := bytes.NewBuffer(jsonByte)
	res, resp, err := replica.client.Post("http://localhost:3000/api/replicac/create/", byts, replica.res)
	if err != nil {
		log.Fatal(err)
	}

	replica.Equal(http.StatusOK, resp.StatusCode)
	replica.Equal("success create replicac", res.Message)
}

func (replica *ReplicaControllerTest) TestGetAllReplica() {
	res, resp, err := replica.client.Get("http://localhost:3000/api/replicac/getAll/", replica.res)
	if err != nil {
		log.Fatal(err)
	}

	replica.Equal(http.StatusOK, resp.StatusCode)
	replica.Equal("success get all data", res.Message)
}

func (replica *ReplicaControllerTest) TestGetOneReplica() {
	res, resp, err := replica.client.Get("http://localhost:3000/api/replicac/get/sample-replicac", replica.res)
	if err != nil {
		log.Fatal(err)
	}

	replica.Equal(http.StatusOK, resp.StatusCode)
	replica.Equal("success get replica by this name", res.Message)
}

func (replica *ReplicaControllerTest) TestDeleteReplica() {
	resp, err := replica.client.Delete("http://localhost:3000/api/replicac/delete/sample-replicac", replica.res)
	if err != nil {
		log.Fatal(err)
	}

	replica.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestReplicaController(t *testing.T) {
	suite.Run(t, NewReplicaControllerTest())
}
