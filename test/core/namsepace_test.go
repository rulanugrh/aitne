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

type NamespaceTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewNamespaceTest() *NamespaceTest {
	return &NamespaceTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (namespace *NamespaceTest) TestCreateNamespace() {
	req := model.Namespace{
		Kind:       "Namespace",
		APIVersion: "v1",
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName:  "create-namespace",
			Labels: map[string]string{
				"type": "test-namespace",
			},
			Annotations: map[string]string{
				"type": "test-namespace",
			},
		},
		Finalizer: []string{
			"test-namespace",
		},
	}

	jsonByte, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonByte)
	res, resp, err := namespace.client.Post("http://localhost:3000/api/namespace/create/", byt, namespace.res)
	if err != nil {
		log.Fatal(err)
	}

	namespace.Equal(http.StatusOK, resp.StatusCode)
	namespace.Equal("success create namespace", res.Message)
}

func (namespace *NamespaceTest) TestGetAllNamespace() {
	res, resp, err := namespace.client.Get("http://localhost:3000/api/namespace/getAll/", namespace.res)
	if err != nil {
		log.Fatal(err)
	}

	namespace.Equal(http.StatusOK, resp.StatusCode)
	namespace.Equal("success get all data", res.Message)
}

func (namespace *NamespaceTest) TestGetOneNamespace() {
	res, resp, err := namespace.client.Get("http://localhost:3000/api/namespace/get/create-namespace", namespace.res)
	if err != nil {
		log.Fatal(err)
	}

	namespace.Equal(http.StatusOK, resp.StatusCode)
	namespace.Equal("success get namespace by this name", res.Message)
}

func (namespace *NamespaceTest) TestDeleteNamespace() {
	resp, err := namespace.client.Delete("http://localhost:3000/api/namespace/delete/create-namespace", namespace.res)
	if err != nil {
		log.Fatal(err)
	}

	namespace.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestNamespace(t *testing.T) {
	suite.Run(t, NewNamespaceTest())
}
