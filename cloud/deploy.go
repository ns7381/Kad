package cloud

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/ns7381/Kad/models"
)

func int32Ptr(i int32) *int32 { return &i }
func Deploy(deployment models.Deployment) error {
	deploymentsClient := K8sClient().AppsV1().Deployments(apiv1.NamespaceDefault)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deployment.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deployment.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deployment.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: deployment.ContainerImage,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(deploy)
	if err != nil {
		panic(err)
		return err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return err
}
