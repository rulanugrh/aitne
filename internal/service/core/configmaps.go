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

type ConfigMaps interface {
	Create(req model.ConfigMap) (*model.ResponseConfigMap, error)
	List() (*[]model.ResponseConfigMap, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseConfigMap, error)
}

type configmap struct {
	client v1.ConfigMapInterface
}

func NewConfigMap(client *kubernetes.Clientset) ConfigMaps {
	return &configmap{
		client: client.CoreV1().ConfigMaps(metav1.NamespaceDefault),
	}
}

func (c *configmap) Create(req model.ConfigMap) (*model.ResponseConfigMap, error) {
	_configmap := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: req.APIVersion,
			Kind:       req.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: req.Meta.Annotations,
			Labels:      req.Meta.Labels,
			Name:        req.Meta.MetaName,
			Namespace:   req.Meta.Namespace,
		},
		Immutable:  req.Immutable,
		BinaryData: req.BinaryData,
		Data:       req.Data,
	}

	data, err := c.client.Create(context.TODO(), &_configmap, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create config map")
	}

	response := model.ResponseConfigMap{
		APIVersion: data.APIVersion,
		Kind:       data.Kind,
		Meta: model.ObjectMeta{
			Namespace:   data.Namespace,
			MetaName:    data.Name,
			Annotations: data.Annotations,
			Labels:      data.Labels,
		},
		Immutable:  data.Immutable,
		BinaryData: data.BinaryData,
		Data:       data.Data,
	}

	return &response, nil
}

func (c *configmap) List() (*[]model.ResponseConfigMap, error) {
	list, err := c.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get list data")
	}

	var response []model.ResponseConfigMap
	for _, data := range list.Items {
		result := model.ResponseConfigMap{
			APIVersion: data.APIVersion,
			Kind:       data.Kind,
			Meta: model.ObjectMeta{
				Namespace:   data.Namespace,
				MetaName:    data.Name,
				Annotations: data.Annotations,
				Labels:      data.Labels,
			},
			Immutable:  data.Immutable,
			BinaryData: data.BinaryData,
			Data:       data.Data,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (c *configmap) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := c.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("cannot deleted data with this name")
	}

	return nil
}

func (c *configmap) GetByName(name string) (*model.ResponseConfigMap, error) {
	data, err := c.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get data by this name")
	}

	response := model.ResponseConfigMap{
		APIVersion: data.APIVersion,
		Kind:       data.Kind,
		Meta: model.ObjectMeta{
			Namespace:   data.Namespace,
			MetaName:    data.Name,
			Annotations: data.Annotations,
			Labels:      data.Labels,
		},
		Immutable:  data.Immutable,
		BinaryData: data.BinaryData,
		Data:       data.Data,
	}

	return &response, nil
}
