// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContainerOption func(c *corev1.Container)

type Resource struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type ResourceRequire struct {
	Limit   Resource
	Require Resource
}

type Image struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func NewContainer(option ...ContainerOption) *corev1.Container {
	var con = &corev1.Container{}

	for _, v := range option {
		v(con)
	}
	return con
}

func ContainerWithResource(require ResourceRequire) ContainerOption {
	return func(c *corev1.Container) {
		c.Resources = corev1.ResourceRequirements{
			Limits: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse(require.Limit.Cpu),
				corev1.ResourceMemory: resource.MustParse(require.Limit.Memory),
			},
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse(require.Require.Cpu),
				corev1.ResourceMemory: resource.MustParse(require.Require.Memory),
			},
		}
	}
}

func ContainerWithImage(image Image) ContainerOption {
	return func(c *corev1.Container) {
		c.Name = image.Name
		c.Image = image.ImageURL
		c.ImagePullPolicy = corev1.PullAlways
	}
}

func ContainerWithCommand(command, args []string) ContainerOption {
	return func(c *corev1.Container) {
		c.Command = command
		c.Args = args
	}
}

func ContainerWithWorkDir(dir string) ContainerOption {
	return func(c *corev1.Container) {
		c.WorkingDir = dir
	}
}

func ContainerWithEnv(env map[string]string) ContainerOption {
	return func(c *corev1.Container) {
		for k, v := range env {
			var e = corev1.EnvVar{
				Name:  k,
				Value: v,
			}

			c.Env = append(c.Env, e)
		}

	}
}

func ContainerWithVolumeMount(mounts []corev1.VolumeMount) ContainerOption {
	return func(c *corev1.Container) {
		c.VolumeMounts = mounts
	}
}

func ContainerWithPort(ports []corev1.ContainerPort) ContainerOption {
	return func(c *corev1.Container) {
		c.Ports = ports
	}
}

func (k *KubernetesClient) GetPod(ctx context.Context, name string) (*corev1.Pod, error) {
	return k.client.CoreV1().Pods(k.namespace).Get(ctx, name, metav1.GetOptions{})
}

func (k *KubernetesClient) ListPods(ctx context.Context, option metav1.ListOptions) (*corev1.PodList, error) {
	return k.client.CoreV1().Pods(k.namespace).List(ctx, option)
}
