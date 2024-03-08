package service

import (
	"context"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/util"
	apiv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type StatefullSet interface {
	Create(req model.StatefullSet) (*model.ResponseStatefullSet, error)
	List() (*[]model.GetStatefulSet, error)
	GetByName(name string) (*model.GetStatefulSet, error)
	Delete(name string) error
}

type statefullset struct {
	client v1.StatefulSetInterface
}

func NewStatefulSet(client *kubernetes.Clientset) StatefullSet {
	return &statefullset{client: client.AppsV1().StatefulSets(corev1.NamespaceDefault)}
}

func (s *statefullset) Create(req model.StatefullSet) (*model.ResponseStatefullSet, error) {
	statefull := apiv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: req.APIVersion,
			Kind:       req.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.Meta.MetaName,
			Namespace:   req.Meta.Namespace,
			Annotations: req.Meta.Annotations,
			Labels:      req.Meta.Labels,
		},
		Spec: apiv1.StatefulSetSpec{
			Replicas:    &req.Replica,
			ServiceName: req.ServiceName,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        req.Meta.MetaName,
					Namespace:   req.Meta.Namespace,
					Labels:      req.Meta.Labels,
					Annotations: req.Meta.Annotations,
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
			},
		},
	}

	data, err := s.client.Create(context.TODO(), &statefull, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("error cannot create statefull set")
	}

	response := model.ResponseStatefullSet{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Meta: model.ObjectMeta{
			Annotations: data.ObjectMeta.Annotations,
			Labels:      data.ObjectMeta.Labels,
			MetaName:    data.ObjectMeta.Name,
			Namespace:   data.ObjectMeta.Namespace,
		},
	}

	return &response, nil
}

func (s *statefullset) List() (*[]model.GetStatefulSet, error) {
	list, err := s.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get data statefull")
	}

	var response []model.GetStatefulSet
	for _, data := range list.Items {
		result := model.GetStatefulSet{
			NodeName:        data.Spec.Template.Spec.NodeName,
			Name:            data.Name,
			Kind:            data.Kind,
			APIVersions:     data.APIVersion,
			Labels:          data.Labels,
			Annotations:     data.Annotations,
			Replica:         *data.Spec.Replicas,
			ResourceVersion: data.ResourceVersion,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (s *statefullset) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := s.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("cannot deleted data by name")
	}

	return nil
}

func (s *statefullset) GetByName(name string) (*model.GetStatefulSet, error) {
	data, err := s.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot find statefull by this name statefull")
	}

	response := model.GetStatefulSet{
		NodeName:        data.Spec.Template.Spec.NodeName,
		Name:            data.Name,
		Namespace:       data.Namespace,
		Kind:            data.Kind,
		APIVersions:     data.APIVersion,
		Annotations:     data.Annotations,
		Labels:          data.Labels,
		Replica:         *data.Spec.Replicas,
		ResourceVersion: data.ResourceVersion,
	}

	return &response, nil
}
