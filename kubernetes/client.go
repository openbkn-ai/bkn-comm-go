// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	client    *kubernetes.Clientset
	namespace string //命名空间
}

func NewKubernetesClient(namespace string) (*KubernetesClient, error) {
	// config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	config, err := clientcmd.BuildConfigFromFlags("", "")

	if err != nil {
		return nil, err
	}
	cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &KubernetesClient{
		client:    cli,
		namespace: namespace,
	}, nil
}

//func ResourceExist(err error) bool {
//	return err.(*k8serr.StatusError).Status().Reason == v1.StatusReasonAlreadyExists
//}

func ResourceNotFound(err error) bool {

	switch err := err.(type) {
	case *k8serr.StatusError:
		return err.Status().Reason == metav1.StatusReasonNotFound
	}
	return false
}
