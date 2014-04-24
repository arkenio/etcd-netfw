package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := parseConfig()

	//Exit gracefully
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		log.Println("Shutting down...")
		os.Exit(0)
	}()

	b := NewBackends(c)
	b.Init()
	log.Println("Starting proxy")
	p := NewTCPProxy(c, b)
	p.start()

}
