package core

import (
	"context"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Node interface {
	Create(req model.Node) (*model.ResponseNode, error)
	List() (*[]model.ResponseNode, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseNode, error)
}

type node struct {
	client v1.NodeInterface
}

func NewNodeConfig(client *kubernetes.Clientset) Node {
	return &node{
		client: client.CoreV1().Nodes(),
	}
}

func (nd *node) Create(req model.Node) (*model.ResponseNode, error) {
	_node := corev1.Node{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.MetaName,
			Namespace:   req.Namespace,
			Annotations: req.Annotations,
			Labels:      req.Labels,
		},
		Spec: corev1.NodeSpec{
			PodCIDR:    req.PodCIDR,
			ProviderID: req.ProviderID,
			Taints: []corev1.Taint{
				{
					Key:   req.Taint.Key,
					Value: req.Taint.Value,
				},
			},
		},
	}

	data, err := nd.client.Create(context.TODO(), &_node, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create node")
	}

	response := model.ResponseNode{
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		Namespace:   data.Namespace,
		MetaName:    data.Name,
		Labels:      data.Labels,
		Annotations: data.Annotations,
		PodCIDR:     data.Spec.PodCIDR,
		ProviderID:  data.Spec.ProviderID,
	}

	return &response, nil
}

func (nd *node) List() (*[]model.ResponseNode, error) {
	list, err := nd.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get all data node")
	}

	var response []model.ResponseNode
	for _, data := range list.Items {
		result := model.ResponseNode{
			Kind:        data.Kind,
			APIVersion:  data.APIVersion,
			Namespace:   data.Namespace,
			MetaName:    data.Name,
			Labels:      data.Labels,
			Annotations: data.Annotations,
			PodCIDR:     data.Spec.PodCIDR,
			ProviderID:  data.Spec.ProviderID,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (nd *node) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := nd.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("cannot deleted data with this name")
	}

	return nil
}

func (nd *node) GetByName(name string) (*model.ResponseNode, error) {
	data, err := nd.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get data by this name")
	}

	response := model.ResponseNode{
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		Namespace:   data.Namespace,
		MetaName:    data.Name,
		Labels:      data.Labels,
		Annotations: data.Annotations,
		PodCIDR:     data.Spec.PodCIDR,
		ProviderID:  data.Spec.ProviderID,
	}

	return &response, nil
}
