package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/coreos/go-etcd/etcd"
)

func Test_cluster(t *testing.T) {
	var b *backends

	Convey("Given a backend", t, func() {
		b = &backends{}

		Convey("Then the number of instances is 0", func() {
			So(len(b.hosts), ShouldEqual, 0)
		})

		Convey("When i add a node", func() {
			n := &etcd.Node{Key: "1", Value: "{\"host\":\"\",\"port\":8080}"}
			b.Update(n, "add")

			Convey("Then the number of instances is 1", func() {
				So(len(b.hosts), ShouldEqual, 1)

			})

		})

		Convey("When i add several thime the same node", func() {
			n := &etcd.Node{Key: "1", Value: "{\"host\":\"localhost\",\"port\":8080}"}
			b.Update(n, "add")
			b.Update(n, "add")

			Convey("Then the number of instances is 1", func() {
				So(len(b.hosts), ShouldEqual, 1)
				So(b.Next(), ShouldEqual, "localhost:8080")

			})

		})

		Convey("When i add an update node (same key, different value)", func() {
			n1 := &etcd.Node{Key: "1", Value: "{\"host\":\"localhost\",\"port\":8080}"}
			n2 := &etcd.Node{Key: "1", Value: "{\"host\":\"127.0.0.1\",\"port\":8080}"}
			b.Update(n1, "add")
			b.Update(n2, "add")

			Convey("Then the value is updated", func() {
				So(len(b.hosts), ShouldEqual, 1)
				So(b.Next(), ShouldEqual, "127.0.0.1:8080")
			})

		})

	})

}
