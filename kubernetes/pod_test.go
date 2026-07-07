// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	corev1 "k8s.io/api/core/v1"
)

func TestContainerWithResource(t *testing.T) {
	Convey("TestContainerOption", t, func() {
		var dep = NewContainer()
		Convey("test ContainerWithResource", func() {
			var r = ResourceRequire{
				Limit: Resource{
					Cpu:    "2",
					Memory: "2",
				},
				Require: Resource{
					Cpu:    "2",
					Memory: "2",
				},
			}
			ContainerWithResource(r)(dep)
			So(len(dep.Resources.Limits), ShouldEqual, 2)

		})
		Convey("test ContainerWithImage", func() {
			var image = Image{
				Name:     "xx",
				ImageURL: "xx",
			}
			ContainerWithImage(image)(dep)
			So(dep.Image, ShouldEqual, image.ImageURL)
			So(dep.Name, ShouldEqual, image.Name)

		})
		Convey("test ContainerWithCommand", func() {
			cmd := []string{"123"}
			arg := []string{"123"}
			ContainerWithCommand(cmd, arg)(dep)
			So(dep.Command, ShouldResemble, cmd)
			So(dep.Args, ShouldResemble, arg)

		})
		Convey("test ContainerWithWorkDir", func() {
			dir := "dxxx"
			ContainerWithWorkDir(dir)(dep)
			So(dep.WorkingDir, ShouldEqual, dir)

		})
		Convey("test ContainerWithEnv", func() {
			env := map[string]string{
				"x": "x",
			}
			ContainerWithEnv(env)(dep)
			So(len(dep.Env), ShouldEqual, 1)

		})
		Convey("test ContainerWithPort", func() {
			port := []corev1.ContainerPort{
				{
					Name: "xx",
				},
			}
			ContainerWithPort(port)(dep)
			So(len(dep.Ports), ShouldEqual, 1)

		})
		Convey("test ContainerWithVolumeMount", func() {
			mount := []corev1.VolumeMount{
				{
					Name:      "test",
					MountPath: "test",
				},
			}
			ContainerWithVolumeMount(mount)(dep)
			So(len(dep.VolumeMounts), ShouldEqual, 1)

		})

	})
}
