// Copyright The kweaver.ai Authors.
//
// Licensed under the Apache License, Version 2.0.
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeployOption func(deployment *appsv1.Deployment)

func WithContainer(container corev1.Container) DeployOption {
	return func(deployment *appsv1.Deployment) {

		deployment.Spec.Template.Spec.Containers = append(deployment.Spec.Template.Spec.Containers, container)
	}
}

func WithLabel(label map[string]string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.SetLabels(label)
	}
}

func WithAnnotations(annotation map[string]string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.SetAnnotations(annotation)
	}
}

func WithName(name string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Name = name
	}
}

func withNamespace(namespace string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Namespace = namespace
	}
}

func WithMatchLabel(label map[string]string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Selector = metav1.SetAsLabelSelector(label)
	}
}

func WithHostNetwork(is bool) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Template.Spec.HostNetwork = is

	}
}

func WithTemplateAnnotation(anno map[string]string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Template.Annotations = anno
	}
}

func WithServiceAccount(name string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Template.Spec.ServiceAccountName = name
	}
}

func WithPodLabel(label map[string]string) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Template.SetLabels(label)
	}
}

func WithReplicas(i int32) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Replicas = &i
	}
}
func WithStrategy(strategy appsv1.DeploymentStrategy) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.Spec.Strategy = strategy
	}
}
func WithOwnerReference(own []metav1.OwnerReference) DeployOption {
	return func(deployment *appsv1.Deployment) {
		deployment.SetOwnerReferences(own)
	}
}

func newDeploy(volumnMounts []corev1.VolumeMount) *appsv1.Deployment {
	var volumes []corev1.Volume
	for _, val := range volumnMounts {
		volume := corev1.Volume{
			Name: val.Name,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: val.Name,
					},
				},
			},
		}
		volumes = append(volumes, volume)
	}

	var dep = &appsv1.Deployment{
		// 基本属性
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: corev1.PodSpec{
					Volumes: volumes,
				},
			},
		},
	}

	return dep
}

func (k *KubernetesClient) CreateDeploy(ctx context.Context, volumes []corev1.VolumeMount, option ...DeployOption) (*appsv1.Deployment, error) {

	option = append(option, withNamespace(k.namespace))
	deploy := newDeploy(volumes)
	for _, v := range option {
		v(deploy)
	}
	return k.client.AppsV1().Deployments(k.namespace).Create(ctx, deploy, metav1.CreateOptions{})
}

func (k *KubernetesClient) GetDeploy(ctx context.Context, name string) (*appsv1.Deployment, error) {
	return k.client.AppsV1().Deployments(k.namespace).Get(ctx, name, metav1.GetOptions{})
}

func (k *KubernetesClient) ListDeploy(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	return k.client.AppsV1().Deployments(k.namespace).List(ctx, opts)
}

func (k *KubernetesClient) UpdateDeploy(ctx context.Context, dep *appsv1.Deployment) (*appsv1.Deployment, error) {
	return k.client.AppsV1().Deployments(k.namespace).Update(ctx, dep, metav1.UpdateOptions{})
}

func (k *KubernetesClient) DeleteDeploy(ctx context.Context, name string) error {
	return k.client.AppsV1().Deployments(k.namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
