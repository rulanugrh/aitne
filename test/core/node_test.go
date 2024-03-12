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

type NodeTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewNodeTest() *NodeTest {
	return &NodeTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (node *NodeTest) TestCreateNode() {
	req := model.Node{
		APIVersion: "v1",
		Kind:       "Node",
		MetaName:   "sample-node",
		Namespace:  "default",
		Annotations: map[string]string{
			"type": "node",
		},
		Labels: map[string]string{
			"type": "node",
		},
		PodCIDR:    "172.16.10.0/20",
		ProviderID: "sample-node",
		Taint: model.Taint{
			Key:   "key1",
			Value: "value1",
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := node.client.Post("http://localhost:3000/api/node/create/", byt, node.res)
	if err != nil {
		log.Fatal(err)
	}

	node.Equal(http.StatusOK, resp.StatusCode)
	node.Equal("success create node", res.Message)
}

func (node *NodeTest) TestGetAllNode() {
	res, resp, err := node.client.Get("http://localhost:3000/api/node/getAll/", node.res)
	if err != nil {
		log.Fatal(err)
	}

	node.Equal(http.StatusOK, resp.StatusCode)
	node.Equal("success get all data", res.Message)
}

func (node *NodeTest) TestGetOneNode() {
	res, resp, err := node.client.Get("http://localhost:3000/api/node/get/sample-node", node.res)
	if err != nil {
		log.Fatal(err)
	}

	node.Equal(http.StatusOK, resp.StatusCode)
	node.Equal("success get node by this name", res.Message)
}

func (node *NodeTest) TestDeleteNode() {
	resp, err := node.client.Delete("http://localhost:3000/api/node/delete/sample-node", node.res)
	if err != nil {
		log.Fatal(err)
	}

	node.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestNode(t *testing.T) {
	suite.Run(t, NewNodeTest())
}
