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

type DaemonEndpoint interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

type daemon struct {
	service apps.DaemonSet
}

func NewDameonEndpoint(service apps.DaemonSet) DaemonEndpoint {
	return &daemon{
		service: service,
	}
}

func (d *daemon) Create(w http.ResponseWriter, r *http.Request) {
	var req model.DaemonSet
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

	response, err := json.Marshal(constant.Success("sucessfull create daemon", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (d *daemon) Get(w http.ResponseWriter, r *http.Request) {
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

func (d *daemon) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/daemon/delete/")
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

	response, err := json.Marshal(constant.Deleted("successfull delete daemon: ", name))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	w.Write(response)
}

func (d *daemon) GetByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/daemon/get/")
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
