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

type Namespace interface {
	Create(req model.Namespace) (*model.ResponseNamespace, error)
	List() (*[]model.ResponseNamespace, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseNamespace, error)
}

type namespace struct {
	client v1.NamespaceInterface
}

func NewNamespace(client *kubernetes.Clientset) Namespace {
	return &namespace{
		client: client.CoreV1().Namespaces(),
	}
}

func (n *namespace) Create(req model.Namespace) (*model.ResponseNamespace, error) {
	var finalizer []corev1.FinalizerName
	for _, dt := range req.Finalizer {
		fn := dt
		finalizer = append(finalizer, corev1.FinalizerName(fn))
	}

	_namespace := corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.Meta.MetaName,
			Labels:      req.Meta.Labels,
			Namespace:   req.Meta.Namespace,
			Annotations: req.Meta.Annotations,
		},
		Spec: corev1.NamespaceSpec{
			Finalizers: finalizer,
		},
	}

	data, err := n.client.Create(context.TODO(), &_namespace, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create namespace k8s")
	}

	response := model.ResponseNamespace{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Meta: model.ObjectMeta{
			Namespace:   data.ObjectMeta.Namespace,
			MetaName:    data.ObjectMeta.Name,
			Labels:      data.ObjectMeta.Labels,
			Annotations: data.ObjectMeta.Annotations,
		},
	}

	return &response, nil
}

func (n *namespace) List() (*[]model.ResponseNamespace, error) {
	list, err := n.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get data namespace")
	}

	var response []model.ResponseNamespace
	for _, data := range list.Items {
		result := model.ResponseNamespace{
			Kind:       data.Kind,
			APIVersion: data.APIVersion,
			Meta: model.ObjectMeta{
				Annotations: data.Annotations,
				Labels:      data.Labels,
				MetaName:    data.Name,
				Namespace:   data.Namespace,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (n *namespace) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := n.client.Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deleted})
	if err != nil {
		return util.Error("cannot delete namespace by this name")
	}

	return nil
}

func (n *namespace) GetByName(name string) (*model.ResponseNamespace, error) {
	data, err := n.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get namespace by this name")
	}

	result := model.ResponseNamespace{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Meta: model.ObjectMeta{
			Annotations: data.Annotations,
			Labels:      data.Labels,
			MetaName:    data.Name,
			Namespace:   data.Namespace,
		},
	}

	return &result, nil

}
