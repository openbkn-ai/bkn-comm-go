// Copyright The kweaver.ai Authors. 
// 
// Licensed under the Apache License, Version 2.0. 
// See the LICENSE file in the project root for details.

package kubernetes

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestResourceNotFound(t *testing.T) {
	Convey("TestResourceNotFound", t, func() {
		Convey("test common err", func() {
			err := fmt.Errorf("err")
			b := ResourceNotFound(err)
			So(b, ShouldEqual, false)
		})

		Convey("test not found err", func() {
			err := k8serr.NewNotFound(schema.GroupResource{}, "")
			b := ResourceNotFound(err)
			So(b, ShouldEqual, true)
		})
	})
}
