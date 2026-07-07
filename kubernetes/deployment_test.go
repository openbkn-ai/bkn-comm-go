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

func TestDeployOption(t *testing.T) {
	Convey("TestDeployOption", t, func() {
		var dep = newDeploy([]corev1.VolumeMount{})
		Convey("test with container", func() {
			con := corev1.Container{Name: "123"}
			WithContainer(con)(dep)
			So(dep.Spec.Template.Spec.Containers[0].Name, ShouldEqual, con.Name)

		})
		Convey("test WithLabel", func() {
			label := map[string]string{"123": "123"}
			WithLabel(label)(dep)
			So(dep.Labels, ShouldEqual, label)

		})
		Convey("test WithAnnotations", func() {
			label := map[string]string{"123": "123"}
			WithAnnotations(label)(dep)
			So(dep.Annotations, ShouldEqual, label)

		})
		Convey("test WithName", func() {
			name := "123"
			WithName(name)(dep)
			So(dep.Name, ShouldEqual, name)

		})
		Convey("test withNamespace", func() {
			name := "123"
			withNamespace(name)(dep)
			So(dep.Namespace, ShouldEqual, name)

		})
		Convey("test WithMatchLabel", func() {
			label := map[string]string{"123": "123"}
			WithMatchLabel(label)(dep)
			So(dep.Spec.Selector.MatchLabels, ShouldResemble, label)

		})
		Convey("test WithHostNetwork", func() {
			WithHostNetwork(true)(dep)
			So(dep.Spec.Template.Spec.HostNetwork, ShouldEqual, true)

		})
		Convey("test WithTemplateAnnotation", func() {
			WithTemplateAnnotation(map[string]string{"2": "3"})(dep)
			So(len(dep.Spec.Template.Annotations), ShouldEqual, 1)

		})
		Convey("test WithPodLabel", func() {
			label := map[string]string{"123": "123"}
			WithPodLabel(label)(dep)
			So(dep.Spec.Template.Labels, ShouldEqual, label)

		})
		Convey("test WithServiceAccount", func() {
			name := "123"
			WithServiceAccount(name)(dep)
			So(dep.Spec.Template.Spec.ServiceAccountName, ShouldEqual, name)

		})
		Convey("test WithReplicas", func() {
			var i int32 = 1
			WithReplicas(i)(dep)
			So(*dep.Spec.Replicas, ShouldEqual, i)

		})
		Convey("test WithOwnerReference", func() {
			var ref = metav1.OwnerReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       "feed-ingest-executor",
			}
			WithOwnerReference([]metav1.OwnerReference{ref})(dep)
			So(dep.ObjectMeta.OwnerReferences[0].Kind, ShouldEqual, ref.Kind)

		})

	})
}
