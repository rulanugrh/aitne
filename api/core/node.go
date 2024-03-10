package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/service/core"
	"github.com/rulanugrh/aitne/internal/util/constant"
)

type NodeEndpoint interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

type node struct {
	service core.Node
}

func NewNodeEndpoint(service core.Node) NodeEndpoint {
	return &node{
		service: service,
	}
}

func (n *node) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Node
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	json.Unmarshal(body, &req)

	data, err := n.service.Create(req)
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

	response, err := json.Marshal(constant.Success("sucessfull create node", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (n *node) Get(w http.ResponseWriter, r *http.Request) {
	data, err := n.service.List()
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

func (n *node) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/node/delete/")
	err := n.service.Delete(name)
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

	response, err := json.Marshal(constant.Deleted("successfull delete node: ", name))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	w.Write(response)
}

func (n *node) GetByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/node/get/")
	data, err := n.service.GetByName(name)
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
