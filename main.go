package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/golang/glog"
)

func main() {
	c := parseConfig()

	//Exit gracefully
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		glog.Info("Shutting down...")
		os.Exit(0)
	}()

	b := NewBackends(c)
	b.Init()
	glog.Info("Starting proxy")
	p := NewTCPProxy(c, b)
	p.start()

}
