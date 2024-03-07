package model

type CreateDeployment struct {
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	Name        string            `json:"name"`
	Replicas    int32             `json:"replica"`
	Labels      map[string]string `json:"labels"`
	MatchLabels map[string]string `json:"match_labels"`
	Annotations map[string]string `json:"annotations"`
	Container   Container         `json:"container"`
}

type Container struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	PortExposed  int    `json:"port"`
	NameProtocol string `json:"name_protocol"`
	Protocol     string `json:"protocol"`
}

type ResponseCreate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type GetDeployment struct {
	Name            string            `json:"name"`
	Replica         int32             `json:"replica"`
	Namespace       string            `json:"namespace"`
	APIVersions     string            `json:"api_version"`
	Labels          map[string]string `json:"labels"`
	Annotations     map[string]string `json:"annotations"`
	ResourceVersion string            `json:"resource_v"`
	Kind            string            `json:"kind"`
}
