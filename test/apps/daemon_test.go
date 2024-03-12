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

type DaemonTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewDaemonTest() *DaemonTest {
	return &DaemonTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (daemon *DaemonTest) TestCreateDaemon() {
	req := model.DaemonSet{
		Kind:       "DaemonSet",
		APIVersion: "apps/v1",
		Container: model.Container{
			Name:         "web-nginx",
			Image:        "nginx:latest",
			PortExposed:  80,
			Protocol:     "http",
			NameProtocol: "tcp",
		},
		MatchLabels: map[string]string{
			"type": "web-apps",
		},
		MinReady: 3,
		Meta: model.ObjectMeta{
			Annotations: map[string]string{
				"type": "web-apps",
			},
			Labels: map[string]string{
				"type": "web-apps",
			},
			MetaName:  "web-apps",
			Namespace: "default",
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := daemon.client.Post("http://localhost:3000/api/daemon/create/", byt, daemon.res)
	if err != nil {
		log.Fatal(err)
	}

	daemon.Equal(http.StatusOK, resp.StatusCode)
	daemon.Equal("success create daemon", res.Message)
}

func (daemon *DaemonTest) TestListDaemon() {
	res, resp, err := daemon.client.Get("http://localhost:3000/api/daemon/getAll/", daemon.res)
	if err != nil {
		log.Fatal(err)
	}

	daemon.Equal(http.StatusOK, resp.StatusCode)
	daemon.Equal("success get all data", res.Message)
}

func (daemon *DaemonTest) TestGetOneDaemon() {
	res, resp, err := daemon.client.Get("http://localhost:3000/api/daemon/get/web-apps", daemon.res)
	if err != nil {
		log.Fatal(err)
	}

	daemon.Equal(http.StatusOK, resp.StatusCode)
	daemon.Equal("success get data by this name", res.Message)
}

func (daemon *DaemonTest) TestDeleteDaemon() {
	res, err := daemon.client.Delete("http://localhost:3000/api/daemon/delete/web-apps", daemon.res)
	if err != nil {
		log.Fatal(err)
	}

	daemon.Equal(http.StatusOK, res.StatusCode)
}

func TestDaemon(t *testing.T) {
	suite.Run(t, NewDaemonTest())
}
