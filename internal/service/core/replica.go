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

type ReplicationController interface {
	Create(req model.ReplicationController) (*model.ResponseRC, error)
	List() (*[]model.ResponseRC, error)
	Delete(name string) error
	GetByName(name string) (*model.ResponseRC, error)
}

type replicac struct {
	client v1.ReplicationControllerInterface
}

func NewReplicationController(client *kubernetes.Clientset) ReplicationController {
	return &replicac{
		client: client.CoreV1().ReplicationControllers(metav1.NamespaceDefault),
	}
}

func (r *replicac) Create(req model.ReplicationController) (*model.ResponseRC, error) {
	replica := corev1.ReplicationController{
		TypeMeta: metav1.TypeMeta{
			Kind:       req.Kind,
			APIVersion: req.APIVersion,
		},

		ObjectMeta: metav1.ObjectMeta{
			Annotations: req.Meta.Annotations,
			Labels:      req.Meta.Labels,
			Name:        req.Meta.MetaName,
			Namespace:   req.Meta.Namespace,
		},
		Spec: corev1.ReplicationControllerSpec{
			Replicas:        &req.Replica,
			MinReadySeconds: *&req.MinReady,
			Selector:        req.Selector,
			Template: &corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  req.Container.Name,
							Image: req.Container.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          req.Container.NameProtocol,
									Protocol:      corev1.Protocol(req.Container.Protocol),
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
		return nil, util.Error("sorry cannot create replication controller")
	}

	response := model.ResponseRC{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Selector:   data.Spec.Selector,
		Replica:    data.Spec.Replicas,
		MinReady:   &data.Spec.MinReadySeconds,
		Meta: model.ObjectMeta{
			Namespace:   data.Namespace,
			MetaName:    data.Name,
			Labels:      data.Labels,
			Annotations: data.Annotations,
		},
	}

	return &response, nil
}

func (r *replicac) List() (*[]model.ResponseRC, error) {
	list, err := r.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("cannot get data replication controller")
	}

	var response []model.ResponseRC
	for _, data := range list.Items {
		result := model.ResponseRC{
			Kind:       data.Kind,
			APIVersion: data.APIVersion,
			Selector:   data.Spec.Selector,
			Replica:    data.Spec.Replicas,
			MinReady:   &data.Spec.MinReadySeconds,
			Meta: model.ObjectMeta{
				Namespace:   data.Namespace,
				MetaName:    data.Name,
				Labels:      data.Labels,
				Annotations: data.Annotations,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (r *replicac) Delete(name string) error {
	deleted := metav1.DeletePropagationForeground
	err := r.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deleted,
	})

	if err != nil {
		return util.Error("cannot delete data by this this name")
	}

	return nil
}

func (r *replicac) GetByName(name string) (*model.ResponseRC, error) {
	data, err := r.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("cannot get data by this name")
	}

	response := model.ResponseRC{
		Kind:       data.Kind,
		APIVersion: data.APIVersion,
		Selector:   data.Spec.Selector,
		Replica:    data.Spec.Replicas,
		MinReady:   &data.Spec.MinReadySeconds,
		Meta: model.ObjectMeta{
			Namespace:   data.Namespace,
			MetaName:    data.Name,
			Labels:      data.Labels,
			Annotations: data.Annotations,
		},
	}

	return &response, nil
}
