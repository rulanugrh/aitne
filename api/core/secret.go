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

type SecretEndpoint interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

type secret struct {
	srv core.Secret
}

func NewSecretEndpoint(srv core.Secret) SecretEndpoint {
	return &secret{
		srv: srv,
	}
}

func (s *secret) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Secret
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	json.Unmarshal(body, &req)

	data, err := s.srv.Create(req)
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

	response, err := json.Marshal(constant.Success("sucessfull create secret", data))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (s *secret) Get(w http.ResponseWriter, r *http.Request) {
	data, err := s.srv.List()
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

func (s *secret) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/secret/delete/")
	err := s.srv.Delete(name)
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

	response, err := json.Marshal(constant.Deleted("successfull delete secret: ", name))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	w.Write(response)
}

func (s *secret) GetByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/secret/get/")
	data, err := s.srv.GetByName(name)
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
