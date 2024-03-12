package helper

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rulanugrh/aitne/internal/util/constant"
)

type SuiteInterface interface {
	Result(res *constant.Response, http *http.Response)
	Post(path string, body *bytes.Buffer, res *constant.Response) (*constant.Response, *http.Response, error)
	Get(path string, res *constant.Response) (*constant.Response, *http.Response, error)
	Delete(path string, res *constant.Response) (*http.Response, error)
}

type Suites struct {
	client *http.Client
}

func NewSuiteUtils(client *http.Client) SuiteInterface {
	return &Suites{
		client: client,
	}
}

func (s *Suites) sendRequest(method string, path string, body *bytes.Buffer) (*http.Response, error) {
	bdy := bytes.NewBuffer(nil)
	if body != nil {
		bdy = body
	}

	req, err := http.NewRequest(method, path, bdy)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	return resp, err
}

func (s *Suites) Result(res *constant.Response, http *http.Response) {
	decoder := json.NewDecoder(http.Body)
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
}

func (s *Suites) Post(path string, body *bytes.Buffer, res *constant.Response) (*constant.Response, *http.Response, error) {
	resp, err := s.sendRequest("POST", path, body)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Get(path string, res *constant.Response) (*constant.Response, *http.Response, error) {
	resp, err := s.sendRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Delete(path string, res *constant.Response) (*http.Response, error) {
	resp, err := s.sendRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	s.Result(res, resp)
	return resp, nil
}
