// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSvcOption(t *testing.T) {
	Convey("TestContainerOption", t, func() {
		var dep = newService()
		Convey("test SvcWithLabel", func() {
			label := map[string]string{
				"x": "x",
			}
			SvcWithLabel(label)(dep)
			So(dep.Labels, ShouldEqual, label)

		})
		Convey("test SvcWithAnnotations", func() {
			label := map[string]string{
				"x": "x",
			}
			SvcWithAnnotations(label)(dep)
			So(dep.Annotations, ShouldEqual, label)

		})
		Convey("test SvcWithName", func() {
			var name = "123"
			SvcWithName(name)(dep)
			So(dep.Name, ShouldEqual, name)
		})
		Convey("test svcWithNamespace", func() {
			namespace := "anyrobot"
			svcWithNamespace(namespace)(dep)
			So(dep.Namespace, ShouldEqual, namespace)

		})
		Convey("test SvcWithSpec", func() {
			var spec = corev1.ServiceSpec{
				ExternalName: "x",
			}
			SvcWithSpec(spec)(dep)
			So(dep.Spec.ExternalName, ShouldEqual, spec.ExternalName)

		})

		Convey("test SvcWithOwnerReference", func() {
			var ref = metav1.OwnerReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       "feed-ingest-executor",
			}
			SvcWithOwnerReference([]metav1.OwnerReference{ref})(dep)
			So(dep.ObjectMeta.OwnerReferences[0].Kind, ShouldEqual, ref.Kind)

		})

	})
}
