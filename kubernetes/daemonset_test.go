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

func TestDaemonsetOption(t *testing.T) {
	Convey("TestDsOption", t, func() {
		var ds = newDaemonSet()
		Convey("test with container", func() {
			con := corev1.Container{Name: "123"}
			DsWithContainer(con)(ds)
			So(ds.Spec.Template.Spec.Containers[0].Name, ShouldEqual, con.Name)

		})
		Convey("test WithLabel", func() {
			label := map[string]string{"123": "123"}
			DsWithLabels(label)(ds)
			So(ds.Labels, ShouldEqual, label)

		})

		Convey("test WithName", func() {
			name := "123"
			DsWithName(name)(ds)
			So(ds.Name, ShouldEqual, name)

		})
		Convey("test withNamespace", func() {
			name := "123"
			DsWithNamespace(name)(ds)
			So(ds.Namespace, ShouldEqual, name)

		})
		Convey("test WithMatchLabel", func() {
			label := map[string]string{"123": "123"}
			DsWithMatchLabels(label)(ds)
			So(ds.Spec.Selector.MatchLabels, ShouldResemble, label)

		})
		Convey("test WithPodLabel", func() {
			label := map[string]string{"123": "123"}
			DsWithTemplateLabels(label)(ds)
			So(ds.Spec.Template.Labels, ShouldEqual, label)

		})
		Convey("test WithOwnerReference", func() {
			var ref = metav1.OwnerReference{
				APIVersion: "apps/v1",
				Kind:       "Daemonset",
				Name:       "system-metric-executor",
			}
			DsWithOwnerReference([]metav1.OwnerReference{ref})(ds)
			So(ds.ObjectMeta.OwnerReferences[0].Kind, ShouldEqual, ref.Kind)
		})
		Convey("test WithHostname", func() {
			var hostname = "test"
			DsWithHostname(hostname)(ds)
			So(ds.Spec.Template.Spec.Hostname, ShouldEqual, hostname)
		})
		Convey("test WithHostNetwork", func() {
			DsWithHostNetwork()(ds)
			So(ds.Spec.Template.Spec.HostNetwork, ShouldEqual, true)
		})

	})
}
