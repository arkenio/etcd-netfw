package main

import (
	"flag"
	"log"
)

type Config struct {
	acceptAddr  string
	servicePath string
	etcdAddress string
}

func parseConfig() *Config {
	config := &Config{}
	flag.StringVar(&config.acceptAddr, "a", "0.0.0.0:1337", "local address to listen to")
	flag.StringVar(&config.etcdAddress, "etcdAddress", "http://127.0.0.1:4001/", "etcd client host")
	flag.StringVar(&config.servicePath, "servicePath", "", "key path of the service")
	flag.Parse()

	log.Printf("Dumping Configuration")
	log.Printf("  acceptAddr  : %s", config.acceptAddr)
	log.Printf("  etcdAddress : %s", config.etcdAddress)
	log.Printf("  servicePath : %s", config.servicePath)

	return config
}
