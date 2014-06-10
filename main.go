package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/golang/glog"
	"runtime/pprof"
)

func main() {
	c := parseConfig()

	handleSignals()

	b := NewBackends(c)

	b.Init()
	glog.Info("Starting proxy")
	p := NewTCPProxy(c, b)
	p.start()

}


func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(signals, os.Interrupt, syscall.SIGUSR1)
	go func() {
		sig := <-signals
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			//Exit gracefully
			glog.Info("Shutting down...")
			os.Exit(0)
		case syscall.SIGUSR1:
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		}
	}()

}
