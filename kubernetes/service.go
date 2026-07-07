// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceOption func(service *corev1.Service)

func newService() *corev1.Service {
	svc := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}
	return svc
}

func SvcWithOwnerReference(own []metav1.OwnerReference) ServiceOption {
	return func(svc *corev1.Service) {
		svc.SetOwnerReferences(own)
	}
}

func SvcWithLabel(label map[string]string) ServiceOption {
	return func(svc *corev1.Service) {
		svc.SetLabels(label)
	}
}

func SvcWithAnnotations(annotation map[string]string) ServiceOption {
	return func(svc *corev1.Service) {
		svc.SetAnnotations(annotation)
	}
}

func SvcWithName(name string) ServiceOption {
	return func(svc *corev1.Service) {
		svc.SetName(name)
	}
}

func svcWithNamespace(nameSpace string) ServiceOption {
	return func(svc *corev1.Service) {
		svc.SetNamespace(nameSpace)
	}
}

func SvcWithSpec(spec corev1.ServiceSpec) ServiceOption {
	return func(svc *corev1.Service) {
		svc.Spec = spec
	}
}

func (k *KubernetesClient) CreateSvc(ctx context.Context, option ...ServiceOption) (*corev1.Service, error) {

	option = append(option, svcWithNamespace(k.namespace))
	svc := newService()
	for _, v := range option {
		v(svc)
	}

	return k.client.CoreV1().Services(k.namespace).Create(ctx, svc, metav1.CreateOptions{})
}

func (k *KubernetesClient) DeleteSvc(ctx context.Context, name string) error {
	return k.client.CoreV1().Services(k.namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (k *KubernetesClient) GetSvc(ctx context.Context, name string) (*corev1.Service, error) {
	return k.client.CoreV1().Services(k.namespace).Get(ctx, name, metav1.GetOptions{})
}
