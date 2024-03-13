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

type ConfigMapTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewConfigMapTest() *ConfigMapTest {
	return &ConfigMapTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (config *ConfigMapTest) TestCreateConfigMap() {
	req := model.ConfigMap{
		Kind:       "ConfigMap",
		APIVersion: "v1",
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName:  "sample-configmap",
			Annotations: map[string]string{
				"type": "config",
			},
			Labels: map[string]string{
				"type": "config-map",
			},
		},
		Immutable: true,
		Data: map[string]string{
			"ver_data": "v1.0.0",
			"type.env": "env.dev",
		},
	}

	jsonByte, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byts := bytes.NewBuffer(jsonByte)
	res, resp, err := config.client.Post("http://localhost:3000/api/configmap/create/", byts, config.res)
	if err != nil {
		log.Fatal(err)
	}

	config.Equal(http.StatusOK, resp.StatusCode)
	config.Equal("success create configmap", res.Message)
}

func (config *ConfigMapTest) TestGetAllConfigMap() {
	res, resp, err := config.client.Get("http://localhost:3000/api/configmap/getAll/", config.res)
	if err != nil {
		log.Fatal(err)
	}

	config.Equal(http.StatusOK, resp.StatusCode)
	config.Equal("success get all data", res.Message)
}

func (config *ConfigMapTest) TestGetOneConfigMap() {
	res, resp, err := config.client.Get("http://localhost:3000/api/configmap/get/sample-configmap", config.res)

	if err != nil {
		log.Fatal(err)
	}

	config.Equal(http.StatusOK, resp.StatusCode)
	config.Equal("success get configmap by this name", res.Message)
}

func (config *ConfigMapTest) TestDeleteConfigMap() {
	resp, err := config.client.Delete("http://localhost:3000/api/configmap/delete/sample-configmap", config.res)
	if err != nil {
		log.Fatal(err)
	}

	config.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestConfigMap(t *testing.T) {
	suite.Run(t, NewConfigMapTest())
}
