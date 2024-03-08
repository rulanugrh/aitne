package service

import (
	"context"

	"github.com/rulanugrh/aitne/internal/model"
	"github.com/rulanugrh/aitne/internal/util"
	apiv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type ReplicaSet interface {
	Create(req model.ReplicaSet) (*model.ResponseReplicaSet, error)
	List() (*[]model.GetReplicaSet, error)
	Delete(name string) error
	GetByName(name string) (*model.GetReplicaSet, error)
}

type replicaset struct {
	client v1.ReplicaSetInterface
}

func NewReplicaSet(client v1.ReplicaSetInterface) ReplicaSet {
	return &replicaset{
		client: client,
	}
}

func (r *replicaset) Create(req model.ReplicaSet) (*model.ResponseReplicaSet, error) {
	replica := apiv1.ReplicaSet{
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
		Spec: apiv1.ReplicaSetSpec{
			Replicas: &req.MinReady,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        req.Meta.MetaName,
					Namespace:   req.Meta.Namespace,
					Annotations: req.Meta.Annotations,
					Labels:      req.Meta.Labels,
				},
				Spec: corev1.PodSpec{
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

	data, err := r.client.Create(context.TODO(), &replica, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("error cannot create replica")
	}

	response := model.ResponseReplicaSet{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Meta: model.ObjectMeta{
			Namespace:   data.ObjectMeta.Namespace,
			Labels:      data.ObjectMeta.Labels,
			MetaName:    data.ObjectMeta.Name,
			Annotations: data.ObjectMeta.Annotations,
		},
	}

	return &response, nil
}

func (r *replicaset) List() (*[]model.GetReplicaSet, error) {
	list, err := r.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get data list replica")
	}

	var response []model.GetReplicaSet
	for _, data := range list.Items {
		result := model.GetReplicaSet{
			Name:            data.Name,
			Namespace:       data.Namespace,
			Replica:         *data.Spec.Replicas,
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

func (r *replicaset) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := r.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("error cannot delete replica")
	}

	return nil
}

func (r *replicaset) GetByName(name string) (*model.GetReplicaSet, error) {
	data, err := r.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("error cannot get replica by this name")
	}

	response := model.GetReplicaSet{
		Name:            data.Name,
		Namespace:       data.Namespace,
		Labels:          data.Labels,
		APIVersions:     data.APIVersion,
		Kind:            data.Kind,
		Annotations:     data.Annotations,
		ResourceVersion: data.ResourceVersion,
	}

	return &response, nil
}
