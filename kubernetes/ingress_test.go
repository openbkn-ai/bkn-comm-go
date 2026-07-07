// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIngressOption(t *testing.T) {
	Convey("TestIngressOption", t, func() {
		var dep = newIngress()
		Convey("test IngWithAnnotations", func() {
			label := map[string]string{"123": "123"}
			IngWithAnnotations(label)(dep)
			So(dep.Annotations, ShouldEqual, label)

		})
		Convey("test IngWithName", func() {
			name := "123"
			IngWithName(name)(dep)
			So(dep.Name, ShouldEqual, name)

		})
		Convey("test ingWithNamespace", func() {
			name := "123"
			ingWithNamespace(name)(dep)
			So(dep.Namespace, ShouldEqual, name)

		})
		Convey("test IngWithRule", func() {
			rules := []Rule{
				{
					Path:        "/xx",
					ServiceName: "xx",
					ServicePort: 10086,
				},
			}
			IngWithRule(rules)(dep)
			So(dep.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path, ShouldEqual, rules[0].Path)

		})
		Convey("test IngWithOwnerReference", func() {
			var ref = metav1.OwnerReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       "feed-ingest-executor",
			}
			IngWithOwnerReference([]metav1.OwnerReference{ref})(dep)
			So(dep.ObjectMeta.OwnerReferences[0].Kind, ShouldEqual, ref.Kind)

		})

	})
}
