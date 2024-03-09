package apps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/service/apps"
	"github.com/rulanugrh/aitne/internal/util/constant"
)

type DeploymentEndpoint interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

type deployment struct {
	service apps.Deployment
}

func NewDeploymentEndpoint(service apps.Deployment) DeploymentEndpoint {
	return &deployment{
		service: service,
	}
}

func (d *deployment) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateDeployment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	json.Unmarshal(body, &req)

	data, err := d.service.Create(req)
	if err != nil {
		response, err := json.Marshal(constant.BadRequest("sorry invalid request data"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(constant.Success("sucessfull create deployment", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (d *deployment) Get(w http.ResponseWriter, r *http.Request) {
	data, err := d.service.List()
	if err != nil {
		response, err := json.Marshal(constant.NotFound("sorry data not found"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(404)
		w.Write(response)
		return
	}

	response, err := json.Marshal(constant.Success("successfull get data", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (d *deployment) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/deployment/delete/")
	err := d.service.Delete(name)
	if err != nil {
		response, err := json.Marshal(constant.BadRequest("sorry invalid request data"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(constant.Deleted("successfull delete deployment: ", name))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	w.Write(response)
}

func (d *deployment) GetByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/deployment/get/")
	data, err := d.service.GetByName(name)
	if err != nil {
		response, err := json.Marshal(constant.BadRequest("sorry invalid request data"))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response, err := json.Marshal(constant.Success("successfull get data by this name", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
