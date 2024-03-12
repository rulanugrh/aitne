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

type StatefullSetTest struct {
	suite.Suite
	client helper.SuiteInterface
	res    *constant.Response
}

func NewStatefullSetTest() *StatefullSetTest {
	return &StatefullSetTest{
		client: helper.NewSuiteUtils(&http.Client{}),
		res:    &constant.Response{},
	}
}

func (st *StatefullSetTest) TestCreateStatefull() {
	req := model.StatefullSet{
		Kind:       "StatefullSet",
		APIVersion: "apps/v1",
		Container: model.Container{
			Name:         "demo-statefull",
			Image:        "nginx:1.13",
			PortExposed:  8000,
			Protocol:     "tcp",
			NameProtocol: "http",
		},
		MatchLabels: map[string]string{
			"app": "demo-statefull",
		},
		Meta: model.ObjectMeta{
			Namespace: "default",
			MetaName:  "demo-statefull",
			Annotations: map[string]string{
				"app": "demo-statefull",
			},
			Labels: map[string]string{
				"app": "demo-statefull",
			},
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	res, resp, err := st.client.Post("http://localhost:3000/api/statefull/create/", byt, st.res)
	if err != nil {
		log.Fatal(err)
	}

	st.Equal(http.StatusOK, resp.StatusCode)
	st.Equal("success create statefull", res.Message)
}

func (st *StatefullSetTest) TestGetAllStatefull() {
	res, resp, err := st.client.Get("http://localhost:3000/api/statefull/getAll/", st.res)
	if err != nil {
		log.Fatal(err)
	}

	st.Equal(http.StatusOK, resp.StatusCode)
	st.Equal("success get all data", res.Message)
}

func (st *StatefullSetTest) TestGetOneStatefull() {
	res, resp, err := st.client.Get("http://localhost:3000/api/statefull/get/demo-statefull", st.res)
	if err != nil {
		log.Fatal(err)
	}

	st.Equal(http.StatusOK, resp.StatusCode)
	st.Equal("success get statefull by this name", res.Message)
}

func (st *StatefullSetTest) TestDeleteStatefull() {
	resp, err := st.client.Delete("http://localhost:3000/api/statefull/get/demo-statefull", st.res)
	if err != nil {
		log.Fatal(err)
	}

	st.Equal(http.StatusNoContent, resp.StatusCode)
}
