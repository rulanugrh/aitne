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

type PodTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewPodTest() *PodTest {
	return &PodTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (pod *PodTest) TestCreatePod() {
	req := model.Pod{
		NodeName:   "minikube",
		Kind:       "Pod",
		APIVersion: "v1",
		Container: model.Container{
			Name:         "test-pod",
			Image:        "nginx:latest",
			NameProtocol: "http",
			PortExposed:  80,
			Protocol:     "tcp",
		},
		Replica: 2,
		MatchLabels: map[string]string{
			"app": "test-pod",
		},
		Meta: model.ObjectMeta{
			MetaName:  "test-pod",
			Namespace: "default",
			Annotations: map[string]string{
				"app": "test-pod",
			},
			Labels: map[string]string{
				"app": "test-pod",
			},
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := pod.client.Post("http://localhost:3000/api/pod/create/", byt, pod.res)
	if err != nil {
		log.Fatal(err)
	}

	pod.Equal(http.StatusOK, resp.StatusCode)
	pod.Equal("success create pod", res.Message)
}

func (pod *PodTest) TestGetAllPod() {
	res, resp, err := pod.client.Get("http://localhost:3000/api/pod/getAll/", pod.res)
	if err != nil {
		log.Fatal(err)
	}

	pod.Equal(http.StatusOK, resp.StatusCode)
	pod.Equal("success get all data", res.Message)
}

func (pod *PodTest) TestGetOnePod() {
	res, resp, err := pod.client.Get("http://localhost:3000/api/pod/get/test-pod", pod.res)
	if err != nil {
		log.Fatal(err)
	}

	pod.Equal(http.StatusOK, resp.StatusCode)
	pod.Equal("success get pod by this name", res.Message)
}

func (pod *PodTest) TestDeletePod() {
	resp, err := pod.client.Delete("http://localhost:3000/api/pod/delete/test-pod", pod.res)
	if err != nil {
		log.Fatal(err)
	}

	pod.Equal(http.StatusNoContent, resp.StatusCode)
}

func TestPod(t *testing.T) {
	suite.Run(t, NewPodTest())
}
