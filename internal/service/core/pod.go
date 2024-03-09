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

type Pods interface {
	Create(req model.Pod) (*model.ResponsePod, error)
	List() (*[]model.GetPod, error)
	Delete(name string) error
	GetByName(name string) (*model.GetPod, error)
}

type pod struct {
	client v1.PodInterface
}

func NewPod(client *kubernetes.Clientset) Pods {
	return &pod{
		client: client.CoreV1().Pods(corev1.NamespaceDefault),
	}
}

func (p *pod) Create(req model.Pod) (*model.ResponsePod, error) {
	_pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.Meta.MetaName,
			Namespace:   req.Meta.Namespace,
			Annotations: req.Meta.Annotations,
			Labels:      req.Meta.Labels,
		},
		Spec: corev1.PodSpec{
			NodeName: req.NodeName,
			Containers: []corev1.Container{
				{
					Name:  req.Container.Name,
					Image: req.Container.Image,
					Ports: []corev1.ContainerPort{
						{
							Protocol:      corev1.Protocol(req.Container.Protocol),
							Name:          req.Container.NameProtocol,
							ContainerPort: int32(req.Container.PortExposed),
						},
					},
				},
			},
		},
	}

	data, err := p.client.Create(context.TODO(), &_pod, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create pod")
	}

	response := model.ResponsePod{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Meta: model.ObjectMeta{
			MetaName:    data.Name,
			Namespace:   data.Namespace,
			Labels:      data.Labels,
			Annotations: data.Annotations,
		},
	}

	return &response, nil
}

func (p *pod) List() (*[]model.GetPod, error) {
	list, err := p.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get data all pod")
	}

	var response []model.GetPod
	for _, data := range list.Items {
		result := model.GetPod{
			Name:            data.Name,
			Namespace:       data.Namespace,
			Labels:          data.Labels,
			APIVersions:     data.APIVersion,
			ResourceVersion: data.ResourceVersion,
			Annotations:     data.Annotations,
			Kind:            data.Kind,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (p *pod) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := p.client.Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deleted})
	if err != nil {
		return util.Error("cannot delete pod data by this name")
	}

	return nil
}

func (p *pod) GetByName(name string) (*model.GetPod, error) {
	data, err := p.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get data by this name")
	}

	result := model.GetPod{
		Name:            data.Name,
		Namespace:       data.Namespace,
		Labels:          data.Labels,
		APIVersions:     data.APIVersion,
		ResourceVersion: data.ResourceVersion,
		Annotations:     data.Annotations,
		Kind:            data.Kind,
	}

	return &result, nil
}
