package model

type ConfigMap struct {
	Kind       string            `json:"kind"`
	APIVersion string            `json:"api_version"`
	Data       map[string]string `json:"data"`
	Immutable  bool              `json:"immutable"`
	BinaryData map[string][]byte `json:"binary_data"`
	Meta       ObjectMeta        `json:"meta"`
}

type ResponseConfigMap struct {
	Kind       string            `json:"kind"`
	APIVersion string            `json:"api_version"`
	Meta       ObjectMeta        `json:"meta"`
	Data       map[string]string `json:"data"`
	Immutable  *bool             `json:"immutable"`
	BinaryData map[string][]byte `json:"binary_data"`
}
