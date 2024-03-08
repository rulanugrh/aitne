package model

type Pod struct {
	NodeName    string            `json:"node_name"`
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	Container   Container         `json:"container"`
	Replica     int32             `json:"replica"`
	MatchLabels map[string]string `json:"match_labels"`
	Meta        ObjectMeta        `json:"meta"`
}

type ResponsePod struct {
	Kind       string     `json:"kind"`
	APIVersion string     `json:"api_version"`
	Meta       ObjectMeta `json:"meta"`
}
