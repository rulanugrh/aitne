package model

type ReplicaSet struct {
	Kind        string            `json:"kind"`
	APIVersion  string            `json:"api_version"`
	Container   Container         `json:"container"`
	Replica     int32             `json:"replica"`
	MatchLabels map[string]string `json:"match_labels"`
	Meta        ObjectMeta        `json:"meta"`
}

type ResponseReplicaSet struct {
	Kind       string     `json:"kind"`
	APIVersion string     `json:"api_version"`
	Meta       ObjectMeta `json:"meta"`
}

type GetReplicaSet struct {
	Name            string            `json:"name"`
	Replica         int32             `json:"replica"`
	Namespace       string            `json:"namespace"`
	APIVersions     string            `json:"api_version"`
	Labels          map[string]string `json:"labels"`
	Annotations     map[string]string `json:"annotations"`
	ResourceVersion string            `json:"resource_v"`
	Kind            string            `json:"kind"`
}
