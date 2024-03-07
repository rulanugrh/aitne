package model

type DaemonSet struct {
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	Container   Container         `json:"container"`
	MinReady    int32             `json:"min_ready"`
	MatchLabels map[string]string `json:"match_labels"`
	Meta        ObjectMeta        `json:"meta"`
}

type ObjectMeta struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	MetaName    string            `json:"meta_name"`
	Namespace   string            `json:"namespace"`
}

type ResponseDaemonSet struct {
	Kind       string     `json:"kind"`
	APIVersion string     `json:"api_version"`
	Meta       ObjectMeta `json:"meta"`
}

type GetDaemonSet struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	APIVersions     string            `json:"api_version"`
	Labels          map[string]string `json:"labels"`
	Annotations     map[string]string `json:"annotations"`
	ResourceVersion string            `json:"resource_v"`
	Kind            string            `json:"kind"`
}
