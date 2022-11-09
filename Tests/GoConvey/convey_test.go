package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	convey.Convey("A", t, func() {
		x++
		convey.Convey("A-B", func() {
			x++
			convey.Convey("A-B-C1", func() {

				convey.So(x, convey.ShouldEqual, 2)
			})
			convey.Convey("A-B-C2", func() {
				convey.So(x, convey.ShouldEqual, 4)
			})
			convey.Convey("A-B-C3", func() {
				convey.So(x, convey.ShouldEqual, 6)
			})
		})

		convey.Reset(func() {
			t.Log("finish")
		})
	})
}
