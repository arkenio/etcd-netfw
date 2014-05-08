package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_cluster(t *testing.T) {
	var b *backends

	Convey("Given a backend", t, func() {
		b = &backends{}
		b.config = &Config{servicePath: "/services/postgres"}

		Convey("Then the number of instances is 0", func() {
			So(len(b.hosts), ShouldEqual, 0)
		})


		Convey("When i add a node", func() {
			n := makeLocationNode(b, 1, "", 8080)
			b.Update(n, "add")

			Convey("Then the number of instances is 1", func() {
				So(len(b.hosts), ShouldEqual, 1)

			})

		})

		Convey("When i add several thime the same node", func() {
			n := makeLocationNode(b, 1, "localhost", 8080)
			b.Update(n, "add")
			b.Update(n, "add")

			Convey("Then the number of instances is 1", func() {
				So(len(b.hosts), ShouldEqual, 1)
				So(b.Next(), ShouldEqual, "localhost:8080")

			})

		})

		Convey("When i add an update node (same key, different value)", func() {
			n1 := makeLocationNode(b, 1, "localhost", 8080)
			n2 := makeLocationNode(b, 1, "127.0.0.1", 8080)
			b.Update(n1, "add")
			b.Update(n2, "add")

			Convey("Then the value is updated", func() {
				So(len(b.hosts), ShouldEqual, 1)
				So(b.Next(), ShouldEqual, "127.0.0.1:8080")
			})

		})

	})

}

func makeLocationNode(b *backends, index int, host string, port int) *etcd.Node {
	return &etcd.Node{Key: fmt.Sprintf("%s/%d/location", b.config.servicePath, index),
		Value: fmt.Sprintf("{\"host\":\"%s\",\"port\":%d}", host, port)}

}
