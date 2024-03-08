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

type Deployment interface {
	Create(req model.CreateDeployment) (*model.ResponseCreate, error)
	List() (*[]model.GetDeployment, error)
	Delete(name string) error
	GetByName(name string) (*model.GetDeployment, error)
}

type crudk8s struct {
	client v1.DeploymentInterface
}

func NewDeployment(client *kubernetes.Clientset) Deployment {
	return &crudk8s{
		client: client.AppsV1().Deployments(corev1.NamespaceDefault),
	}
}

func (c *crudk8s) Create(req model.CreateDeployment) (*model.ResponseCreate, error) {
	deploy := apiv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: req.APIVersion,
			Kind:       req.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        req.Name,
			Labels:      req.Labels,
			Annotations: req.Annotations,
		},
		Spec: apiv1.DeploymentSpec{
			Replicas: &req.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: req.MatchLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: req.Labels,
				},
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

	create, err := c.client.Create(context.TODO(), &deploy, metav1.CreateOptions{})
	if err != nil {
		return nil, util.Error("error cannot create deployment")
	}

	response := model.ResponseCreate{
		Name:      create.GetObjectMeta().GetName(),
		Namespace: create.GetObjectMeta().GetNamespace(),
	}

	return &response, nil
}

func (c *crudk8s) List() (*[]model.GetDeployment, error) {
	list, err := c.client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, util.Error("error cannot get data deployment")
	}

	var response []model.GetDeployment
	for _, data := range list.Items {
		result := model.GetDeployment{
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

func (c *crudk8s) Delete(name string) error {
	delete := metav1.DeletePropagationForeground
	err := c.client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &delete,
	})

	if err != nil {
		return util.Error("error cannot get delete deployment by this name")
	}

	return nil
}

func (c *crudk8s) GetByName(name string) (*model.GetDeployment, error) {
	deployment, err := c.client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, util.Error("error cannot get data deployment by this name")
	}

	response := model.GetDeployment{
		Name:            deployment.Name,
		Namespace:       deployment.Namespace,
		Labels:          deployment.Labels,
		APIVersions:     deployment.APIVersion,
		Kind:            deployment.Kind,
		Annotations:     deployment.Annotations,
		Replica:         *deployment.Spec.Replicas,
		ResourceVersion: deployment.ResourceVersion,
	}

	return &response, nil
}
