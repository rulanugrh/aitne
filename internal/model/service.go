package model

type Service struct {
	ServiceType string          `json:"service_type"`
	Kind        string          `json:"kind"`
	APIVersion  string          `json:"api_version"`
	Resource    ResourceService `json:"resources"`
	Replica     int32           `json:"replica"`
	Meta        ObjectMeta      `json:"meta"`
}

type ResourceService struct {
	ExternalIPs    []string          `json:"external_ip"`
	ClusterIP      string            `json:"host_ip"`
	LoadBalancerIP string            `json:"load_balancer"`
	ExternalName   string            `json:"external_name"`
	Ports          Port              `json:"port"`
	Selector       map[string]string `json:"selector"`
}

type Port struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     int32  `json:"port"`
	NodePort int32  `json:"node_port"`
}

type ResponseService struct {
	ServiceType string          `json:"service_type"`
	Kind        string          `json:"kind"`
	APIVersion  string          `json:"api_version"`
	Meta        ObjectMeta      `json:"meta"`
	Resource    ResourceService `json:"resource"`
}
