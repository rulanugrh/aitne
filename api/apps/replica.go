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

type ReplicaEndpoint interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

type replica struct {
	service apps.ReplicaSet
}

func NewReplicaEndpoint(service apps.ReplicaSet) ReplicaEndpoint {
	return &replica{
		service: service,
	}
}

func (rep *replica) Create(w http.ResponseWriter, r *http.Request) {
	var req model.ReplicaSet
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	json.Unmarshal(body, &req)

	data, err := rep.service.Create(req)
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

	response, err := json.Marshal(constant.Success("sucessfull create replica", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (rep *replica) Get(w http.ResponseWriter, r *http.Request) {
	data, err := rep.service.List()
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

func (rep *replica) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/replica/delete/")
	err := rep.service.Delete(name)
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

	response, err := json.Marshal(constant.Deleted("successfull delete replica: ", name))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	w.Write(response)
}

func (rep *replica) GetByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/replica/get/")
	data, err := rep.service.GetByName(name)
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
