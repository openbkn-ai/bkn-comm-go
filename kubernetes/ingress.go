// Copyright The kweaver.ai Authors.
//
// Licensed under the Apache License, Version 2.0.
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"context"

	ingressv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 目前 AR 是v.1.23.4  ingress使用的是 extensions/v1beta,此页所有 api以 AR版本为主进行兼容
// extensions/v1beta1 和 networking.k8s.io/v1beta1 API 版本的 Ingress 不在 v1.22 版本中继续提供
// 迁移清单和 API 客户端使用 networking.k8s.io/v1 API 版本，此 API 从 v1.19 版本开始可用
// 详情请跳转 https://kubernetes.io/zh-cn/docs/reference/using-api/deprecation-guide/

type IngressOption func(ing *ingressv1.Ingress)

func newIngress() *ingressv1.Ingress {
	var ing = &ingressv1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1",
		},
	}
	return ing
}

func IngWithName(name string) IngressOption {
	return func(ing *ingressv1.Ingress) {
		ing.SetName(name)
	}
}

func IngWithOwnerReference(own []metav1.OwnerReference) IngressOption {
	return func(ing *ingressv1.Ingress) {
		ing.SetOwnerReferences(own)
	}
}

func IngWithAnnotations(anno map[string]string) IngressOption {
	return func(ing *ingressv1.Ingress) {
		ing.SetAnnotations(
			anno,
		)
	}
}

func ingWithNamespace(namespace string) IngressOption {
	return func(ing *ingressv1.Ingress) {
		ing.SetNamespace(namespace)
	}
}

func IngWithRule(rules []Rule) IngressOption {
	return func(ing *ingressv1.Ingress) {
		var r = ingressv1.IngressRule{
			IngressRuleValue: ingressv1.IngressRuleValue{
				HTTP: &ingressv1.HTTPIngressRuleValue{
					Paths: []ingressv1.HTTPIngressPath{},
				},
			},
		}
		for _, rule := range rules {

			var (
				HTTPPath ingressv1.HTTPIngressPath
				pathType = ingressv1.PathTypeImplementationSpecific
			)
			HTTPPath.Path = rule.Path
			HTTPPath.PathType = &pathType
			HTTPPath.Backend = ingressv1.IngressBackend{
				Service: &ingressv1.IngressServiceBackend{
					Port: ingressv1.ServiceBackendPort{
						Number: rule.ServicePort,
					},
					Name: rule.ServiceName,
				},
			}
			r.HTTP.Paths = append(r.HTTP.Paths, HTTPPath)
		}
		ing.Spec.Rules = append(ing.Spec.Rules, r)
	}
}

type Rule struct {
	Path        string
	ServiceName string
	ServicePort int32
}

func (k *KubernetesClient) CreateIngress(ctx context.Context, option ...IngressOption) (*ingressv1.Ingress, error) {
	ingress := newIngress()
	option = append(option, ingWithNamespace(k.namespace))

	for _, v := range option {
		v(ingress)
	}

	return k.client.NetworkingV1().Ingresses(k.namespace).Create(ctx, ingress, metav1.CreateOptions{})
}

func (k *KubernetesClient) DeleteIngress(ctx context.Context, name string) error {
	return k.client.NetworkingV1().Ingresses(k.namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (k *KubernetesClient) GetIngress(ctx context.Context, name string) (*ingressv1.Ingress, error) {
	return k.client.NetworkingV1().Ingresses(k.namespace).Get(ctx, name, metav1.GetOptions{})
}
