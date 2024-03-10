package model

type ReplicationController struct {
	APIVersion string            `json:"api_version"`
	Kind       string            `json:"kind"`
	Meta       ObjectMeta        `json:"meta"`
	Replica    *int32            `json:"replica"`
	MinReady   *int32            `json:"min_ready"`
	Selector   map[string]string `json:"selector"`
	Container  Container         `json:"container"`
}

type ResponseRC struct {
	APIVersion string            `json:"api_version"`
	Kind       string            `json:"kind"`
	Selector   map[string]string `json:"selector"`
	Meta       ObjectMeta        `json:"meta"`
	Replica    *int32            `json:"replica"`
	MinReady   *int32            `json:"min_ready"`
}
