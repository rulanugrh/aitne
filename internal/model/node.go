package model

type Node struct {
	APIVersion  string            `json:"api_version"`
	Kind        string            `json:"kind"`
	MetaName    string            `json:"meta_name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	PodCIDR     string            `json:"pod_cidr"`
	ProviderID  string            `json:"provider_Id"`
	Taint       Taint             `json:"taint"`
}

type ResponseNode struct {
	APIVersion  string            `json:"api_version"`
	Kind        string            `json:"kind"`
	MetaName    string            `json:"meta_name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	PodCIDR     string            `json:"pod_cidr"`
	ProviderID  string            `json:"provider_Id"`
}

type Taint struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
