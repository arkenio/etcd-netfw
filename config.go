package main

import (
	"flag"
	"github.com/golang/glog"
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

	glog.Info("Dumping Configuration")
	glog.Infof("  acceptAddr  : %s", config.acceptAddr)
	glog.Infof("  etcdAddress : %s", config.etcdAddress)
	glog.Infof("  servicePath : %s", config.servicePath)

	return config
}
