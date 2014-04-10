package main

import (
	"fmt"
  "net"
  "os"
  "io"
	"log"
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
	log.Printf("Listening on %s : ",p.config.acceptAddr)
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
		fmt.Fprintf(os.Stderr, "remote dial failed: %v\n", err)
		return
	}
	go io.Copy(local, remote)
	go io.Copy(remote, local)
}

func fatal(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "netfwd: %s\n", fmt.Sprintf(s, a))
	os.Exit(2)
}
