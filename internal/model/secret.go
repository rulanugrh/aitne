package model

type Secret struct {
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	MetaName    string            `json:"meta_name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	Immutable   bool              `json:"immutable"`
	Type        string            `json:"type"`
	Data        map[string][]byte `json:"data"`
	StringData  map[string]string `json:"str_data"`
}

type ResponseSecret struct {
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	MetaName    string            `json:"meta_name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	Immutable   bool              `json:"immutable"`
	Type        string            `json:"type"`
	Data        map[string][]byte `json:"data"`
	StringData  map[string]string `json:"str_data"`
}
