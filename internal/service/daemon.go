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

type DaemonSet interface {
	Create(req model.DaemonSet) (*model.ResponseDaemonSet, error)
	List() (*[]model.GetDaemonSet, error)
	Delete(name string) error
	GetByName(name string) (*model.GetDaemonSet, error)
}
type daemonset struct {
	client v1.DaemonSetInterface
}

func NewDaemonSet(client v1.DaemonSetInterface) DaemonSet {
	return &daemonset{
		client: client,
	}
}

func (d *daemonset) Create(req model.DaemonSet) (*model.ResponseDaemonSet, error) {
	daemon := apiv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels:      req.Meta.Labels,
			Name:        req.Meta.MetaName,
			Namespace:   req.Meta.Namespace,
			Annotations: req.Meta.Annotations,
		},
		Spec: apiv1.DaemonSetSpec{
			MinReadySeconds: req.MinReady,
			Selector: &metav1.LabelSelector{
				MatchLabels: req.MatchLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        req.Meta.MetaName,
					Namespace:   req.Meta.Namespace,
					Labels:      req.Meta.Labels,
					Annotations: req.Meta.Annotations,
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

	data, err := d.client.Create(context.TODO(), &daemon, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("error cannot create daemon")
	}

	response := model.ResponseDaemonSet{
		APIVersion: data.APIVersion,
		Kind:       data.Kind,
		Meta: model.ObjectMeta{
			Namespace:   data.ObjectMeta.Namespace,
			MetaName:    data.ObjectMeta.Name,
			Annotations: data.ObjectMeta.Annotations,
			Labels:      data.ObjectMeta.Labels,
		},
	}

	return &response, nil
}

func (d *daemonset) List() (*[]model.GetDaemonSet, error) {
	data, err := d.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("error cannot get data daemon")
	}

	var response []model.GetDaemonSet
	for _, dt := range data.Items {
		result := model.GetDaemonSet{
			Name:            dt.Name,
			Namespace:       dt.Namespace,
			Kind:            dt.Kind,
			Labels:          dt.Labels,
			APIVersions:     dt.APIVersion,
			Annotations:     dt.Annotations,
			ResourceVersion: dt.ResourceVersion,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (d *daemonset) Delete(name string) error {
	delete := metav1.DeletePropagationForeground
	err := d.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &delete,
	})

	if err != nil {
		return util.Error("error cannot delete daemon by this name")
	}

	return nil
}

func (d *daemonset) GetByName(name string) (*model.GetDaemonSet, error) {
	data, err := d.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("error cannot get data daemon by this name ")
	}

	response := model.GetDaemonSet{
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
