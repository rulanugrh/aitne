package model

type Namespace struct {
	Finalizer  []string   `json:"finalizer"`
	Kind       string     `json:"kind"`
	APIVersion string     `json:"api_version"`
	Meta       ObjectMeta `json:"meta"`
}

type ResponseNamespace struct {
	Kind       string     `json:"kind"`
	APIVersion string     `json:"api_version"`
	Meta       ObjectMeta `json:"meta"`
}
