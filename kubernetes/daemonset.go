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

type HostPathVolume struct {
	// volume name
	VolumeName string
	// the path of directory on host
	HostPath string
}

type DsOption func(*appsv1.DaemonSet)

func DsWithName(name string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Name = name
	}
}

func DsWithNamespace(ns string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Namespace = ns
		ds.Spec.Template.Namespace = ns
	}
}

func DsWithLabels(labels map[string]string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.SetLabels(labels)
	}
}

func DsWithMatchLabels(labels map[string]string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Selector.MatchLabels = labels
	}
}

func DsWithTemplateLabels(labels map[string]string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Template.SetLabels(labels)
	}
}

func DsWithContainer(containers ...corev1.Container) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Template.Spec.Containers = containers
	}
}
func DsWithOwnerReference(ownerRefs []metav1.OwnerReference) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.SetOwnerReferences(ownerRefs)
	}
}

func DsWithHostPathVolumes(hostPathVolumes []HostPathVolume) DsOption {
	// set default path type
	pathType := corev1.HostPathType("")
	volumes := []corev1.Volume{}
	for _, v := range hostPathVolumes {
		volume := corev1.Volume{
			Name: v.VolumeName,
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: v.HostPath,
					Type: &pathType,
				},
			},
		}
		volumes = append(volumes, volume)
	}

	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Template.Spec.Volumes = volumes
	}
}

func DsWithHostname(hostname string) DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Template.Spec.Hostname = hostname
	}
}
func DsWithHostNetwork() DsOption {
	return func(ds *appsv1.DaemonSet) {
		ds.Spec.Template.Spec.HostNetwork = true
	}
}

func (k *KubernetesClient) CreateDaemonSet(ctx context.Context, opts ...DsOption) (*appsv1.DaemonSet, error) {

	daemonset := newDaemonSet()
	for _, opt := range opts {
		opt(daemonset)
	}
	return k.client.AppsV1().DaemonSets(k.namespace).Create(ctx, daemonset, metav1.CreateOptions{})
}

func (k *KubernetesClient) GetDaemonSet(ctx context.Context, name string) (*appsv1.DaemonSet, error) {
	return k.client.AppsV1().DaemonSets(k.namespace).Get(ctx, name, metav1.GetOptions{})
}

func (k *KubernetesClient) UpdateDaemonSet(ctx context.Context, ds *appsv1.DaemonSet, opts ...DsOption) (*appsv1.DaemonSet, error) {
	for _, opt := range opts {
		opt(ds)
	}
	return k.client.AppsV1().DaemonSets(k.namespace).Update(ctx, ds, metav1.UpdateOptions{})
}

func (k *KubernetesClient) DeleteDaemonSet(ctx context.Context, name string) error {
	return k.client.AppsV1().DaemonSets(k.namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func newDaemonSet() *appsv1.DaemonSet {
	daemonset := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{},
			},
		},
	}
	return daemonset
}
