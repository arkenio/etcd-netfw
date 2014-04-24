package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"github.com/golang/glog"
)

type tcpproxy struct {
	config   *Config
	backends *backends
}

func NewTCPProxy(c *Config, backends *backends) *tcpproxy {
	return &tcpproxy{c, backends}
}

func (p *tcpproxy) start() {
	local, err := net.Listen("tcp", p.config.acceptAddr)
	glog.Infof("Listening on %s : ", p.config.acceptAddr)
	if local == nil {
		fatal("cannot listen: %v", err)
	}
	for {
		conn, err := local.Accept()
		if conn == nil {
			fatal("accept failed: %v", err)
		}

		remoteAddr := p.backends.Next()
		go forward(conn, remoteAddr)
	}
}

func forward(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if remote == nil {
		glog.Fatalf("remote dial failed: %v\n", err)
		return
	}
	go io.Copy(local, remote)
	go io.Copy(remote, local)
}

func fatal(s string, a ...interface{}) {
	glog.Fatalf("netfwd: %s\n", fmt.Sprintf(s, a))
	os.Exit(2)
}
