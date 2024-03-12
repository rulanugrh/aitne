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

type SecretTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewSecretTest() *SecretTest {
	return &SecretTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (secret *SecretTest) TestCreateService() {
	req := model.Secret{
		Kind:       "Secret",
		APIVersion: "v1",
		MetaName:   "sample-secret",
		Labels: map[string]string{
			"type": "sample-secret",
		},
		Annotations: map[string]string{
			"type": "sample-secret",
		},
		Immutable: true,
		Namespace: "default",
		Type:      "Opaque",
		StringData: map[string]string{
			"type": "sample-secret",
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byts := bytes.NewBuffer(jsonBytes)
	res, resp, err := secret.client.Post("http://localhost:3000/api/secret/create", byts, secret.res)
	if err != nil {
		log.Fatal(err)
	}

	secret.Equal(http.StatusOK, resp.StatusCode)
	secret.Equal("success create secret", res.Message)
}

func (secret *SecretTest) TestGetAllService() {
	res, resp, err := secret.client.Get("http://localhost:3000/api/secret/getAll/", secret.res)
	if err != nil {
		log.Fatal(err)
	}

	secret.Equal(http.StatusOK, resp.StatusCode)
	secret.Equal("success get all data", res.Message)
}

func (secret *SecretTest) TestGetOneSecret() {
	res, resp, err := secret.client.Get("http://localhost:3000/api/secret/get/sample-secret", secret.res)
	if err != nil {
		log.Fatal(err)
	}

	secret.Equal(http.StatusOK, resp.StatusCode)
	secret.Equal("success get secret by this name", res.Message)
}

func (secret *SecretTest) TestDeleteSecret() {
	resp, err := secret.client.Delete("http://localhost:3000/api/secret/get/sample-secret", secret.res)
	if err != nil {
		log.Fatal(err)
	}

	secret.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestSecret(t *testing.T) {
	suite.Run(t, NewSecretTest())
}
