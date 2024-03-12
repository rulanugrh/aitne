package apps

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/util/constant"
	helper "github.com/rulanugrh/aitne/test"
	"github.com/stretchr/testify/suite"
)

type ReplicaSetTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewReplicaSetTest() *ReplicaSetTest {
	return &ReplicaSetTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (rc *ReplicaSetTest) TestCreateReplicaSet() {
	req := model.ReplicaSet{
		Kind:       "ReplicationSet",
		APIVersion: "apps/v1",
		Container: model.Container{
			Name:         "demo-replica",
			Image:        "nginx:latest",
			NameProtocol: "http",
			PortExposed:  80,
			Protocol:     "tcp",
		},
		MatchLabels: map[string]string{
			"app": "demo",
		},
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName:  "demo-repica",
			Annotations: map[string]string{
				"app": "demo",
			},
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Replica: 2,
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := rc.client.Post("http://localhost:3000/api/replica/create", byt, rc.res)
	if err != nil {
		log.Fatal(err)
	}

	rc.Equal(http.StatusOK, resp.StatusCode)
	rc.Equal("success create replica", res.Message)
}

func (rc *ReplicaSetTest) TestGetAllReplicaSet() {
	res, resp, err := rc.client.Get("http://localhost:3000/api/replica/getAll/", rc.res)
	if err != nil {
		log.Fatal(err)
	}

	rc.Equal(http.StatusOK, resp.StatusCode)
	rc.Equal("success get all data", res.Message)
}

func (rc *ReplicaSetTest) TestGetOneReplicaSet() {
	res, resp, err := rc.client.Get("http://localhost:3000/api/replica/get/demo-replica", rc.res)
	if err != nil {
		log.Fatal(err)
	}

	rc.Equal(http.StatusOK, resp.StatusCode)
	rc.Equal("success get replica by this name", res.Message)
}

func (rc *ReplicaSetTest) TestDeleteReplicaSet() {
	resp, err := rc.client.Delete("http://localhost:3000/api/replica/get/demo-replica", rc.res)
	if err != nil {
		log.Fatal(err)
	}

	rc.Equal(http.StatusNoContent, resp.StatusCode)
}
