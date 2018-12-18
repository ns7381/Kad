package cloud

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewService(name string) error {
	serviceClient := K8sClient().CoreV1().Services(apiv1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 8080,
					TargetPort: intstr.IntOrString{Type:0, IntVal:8080, StrVal:""},
				},
			},
			Selector: map[string]string{
				"app": name,
			},
			Type:apiv1.ServiceTypeNodePort,
		},
	}
	_, err := serviceClient.Create(service)
	return err
}