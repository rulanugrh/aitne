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

type Service interface {
	Create(req model.Service) (*model.ResponseService, error)
	List() (*[]model.ResponseService, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseService, error)
}

type service struct {
	client v1.ServiceInterface
}

func NewServiceKurbenetes(client *kubernetes.Clientset) Service {
	return &service{
		client: client.CoreV1().Services(corev1.NamespaceDefault),
	}
}

func (s *service) Create(req model.Service) (*model.ResponseService, error) {
	_service := corev1.Service{
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
		Spec: corev1.ServiceSpec{
			Type:           corev1.ServiceType(req.ServiceType),
			Selector:       req.Resource.Selector,
			ExternalIPs:    req.Resource.ExternalIPs,
			ExternalName:   req.Resource.ExternalName,
			LoadBalancerIP: req.Resource.LoadBalancerIP,
			Ports: []corev1.ServicePort{
				{
					Name:     req.Resource.Ports.Name,
					Protocol: corev1.Protocol(req.Resource.Ports.Protocol),
					Port:     req.Resource.Ports.Port,
					NodePort: req.Resource.Ports.NodePort,
				},
			},
		},
	}

	data, err := s.client.Create(context.TODO(), &_service, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("cannot create service k8s")
	}

	response := model.ResponseService{
		ServiceType: string(data.Spec.Type),
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		Meta: model.ObjectMeta{
			Annotations: data.Annotations,
			Labels:      data.Labels,
			MetaName:    data.Name,
			Namespace:   data.Namespace,
		},
		Resource: model.ResourceService{
			ExternalIPs:    data.Spec.ExternalIPs,
			ExternalName:   data.Spec.ExternalName,
			ClusterIP:      data.Spec.ClusterIP,
			LoadBalancerIP: data.Spec.LoadBalancerIP,
			Selector:       data.Spec.Selector,
		},
	}

	return &response, nil
}

func (s *service) List() (*[]model.ResponseService, error) {
	list, err := s.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get all data services")
	}

	var response []model.ResponseService
	for _, data := range list.Items {
		result := model.ResponseService{
			ServiceType: string(data.Spec.Type),
			Kind:        data.Kind,
			APIVersion:  data.APIVersion,
			Meta: model.ObjectMeta{
				Annotations: data.Annotations,
				Labels:      data.Labels,
				MetaName:    data.Name,
				Namespace:   data.Namespace,
			},
			Resource: model.ResourceService{
				ExternalIPs:    data.Spec.ExternalIPs,
				ExternalName:   data.Spec.ExternalName,
				ClusterIP:      data.Spec.ClusterIP,
				LoadBalancerIP: data.Spec.LoadBalancerIP,
				Selector:       data.Spec.Selector,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (s *service) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := s.client.Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deleted})

	if err != nil {
		return util.Error("cannot delete data by this name")
	}

	return nil
}

func (s *service) GetByName(name string) (*model.ResponseService, error) {
	data, err := s.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get service by this name")
	}

	response := model.ResponseService{
		ServiceType: string(data.Spec.Type),
		Kind:        data.Kind,
		APIVersion:  data.APIVersion,
		Meta: model.ObjectMeta{
			Annotations: data.Annotations,
			Labels:      data.Labels,
			MetaName:    data.Name,
			Namespace:   data.Namespace,
		},
		Resource: model.ResourceService{
			ExternalIPs:    data.Spec.ExternalIPs,
			ExternalName:   data.Spec.ExternalName,
			ClusterIP:      data.Spec.ClusterIP,
			LoadBalancerIP: data.Spec.LoadBalancerIP,
			Selector:       data.Spec.Selector,
		},
	}

	return &response, nil
}
