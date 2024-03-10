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

type Secret interface {
	Create(req model.Secret) (*model.ResponseSecret, error)
	List() (*[]model.ResponseSecret, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseSecret, error)
}

type secret struct {
	client v1.SecretInterface
}

func NewSecretKurbenetes(client *kubernetes.Clientset) Secret {
	return &secret{
		client: client.CoreV1().Secrets(metav1.NamespaceDefault),
	}
}

func (se *secret) Create(req model.Secret) (*model.ResponseSecret, error) {
	_secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.MetaName,
			Namespace:   req.Namespace,
			Labels:      req.Labels,
			Annotations: req.Annotations,
		},
		Type:       corev1.SecretType(req.Type),
		Data:       req.Data,
		StringData: req.StringData,
		Immutable:  &req.Immutable,
	}

	data, err := se.client.Create(context.TODO(), &_secret, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create secret")
	}

	response := model.ResponseSecret{
		Namespace:   data.Namespace,
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		MetaName:    data.Name,
		Labels:      data.Labels,
		Annotations: data.Annotations,
		Data:        data.Data,
		Immutable:   *data.Immutable,
		Type:        string(data.Type),
		StringData:  data.StringData,
	}

	return &response, nil
}

func (se *secret) List() (*[]model.ResponseSecret, error) {
	list, err := se.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get all data")
	}

	var response []model.ResponseSecret
	for _, data := range list.Items {
		result := model.ResponseSecret{
			Namespace:   data.Namespace,
			Kind:        data.Kind,
			APIVersion:  data.APIVersion,
			MetaName:    data.Name,
			Labels:      data.Labels,
			Annotations: data.Annotations,
			Data:        data.Data,
			Immutable:   *data.Immutable,
			Type:        string(data.Type),
			StringData:  data.StringData,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (se *secret) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := se.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("cannot delete secret by this name")
	}

	return nil
}

func (se *secret) GetByName(name string) (*model.ResponseSecret, error) {
	data, err := se.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get data by this name")
	}

	response := model.ResponseSecret{
		Namespace:   data.Namespace,
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		MetaName:    data.Name,
		Labels:      data.Labels,
		Annotations: data.Annotations,
		Data:        data.Data,
		Immutable:   *data.Immutable,
		Type:        string(data.Type),
		StringData:  data.StringData,
	}

	return &response, nil
}
