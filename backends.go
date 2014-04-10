package main

import (
	"encoding/json"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net"
	"strconv"
	"sync"
)

type service struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type host struct {
	key  string
	addr string
}

type backends struct {
	config     *Config
	path       string
	hosts      []host
	lastIndex  int
	watchIndex uint64
	lock       sync.RWMutex
}

func NewBackends(c *Config) *backends {
  return &backends{config: c}
}

func (b *backends) Init() {
  peers := trimsplit(b.config.etcdAddress,",")
  client := etcd.NewClient(peers)
  b.Sync(client)
  go b.Watch(client)
}

func (b *backends) Dump(action string) {
	for _, v := range b.hosts {
		log.Printf("Dump after %s %s -> %s", action, v.key, v.addr)
	}
}

func (b *backends) Remove(key string) {
	match := -1
	for k, v := range b.hosts {
		if v.key == key {
			match = k
		}
	}

	b.hosts = append(b.hosts[:match], b.hosts[match+1:]...)
	b.Dump("remove")
}

func (b *backends) Update(node *etcd.Node, action string) {
	b.lock.Lock()
	defer b.lock.Unlock()

	log.Printf("key: %s action: %s value: %s", node.Key, action, string(node.Value))

	s := &service{}
	if action == "delete" || action == "expire" {
		b.Remove(node.Key)
		return
	}

	err := json.Unmarshal([]byte(node.Value), s)
	if err != nil {
		panic(err)
	}

	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))

	// TODO: create an actual set data structure
	for _, v := range b.hosts {
		if v.key == node.Key {
			b.Dump(action)
			return
		}
	}

	b.hosts = append(b.hosts, host{addr: addr, key: node.Key})
}

func (b *backends) Watch(client *etcd.Client) {
	receiver := make(chan *etcd.Response)
	go client.Watch(b.config.servicePath, uint64(b.watchIndex), true, receiver, nil)

	for {
		resp := <-receiver
		b.Update(resp.Node, resp.Action)
	}
}

func (b *backends) Sync(client *etcd.Client) error {
  log.Printf("Synchronizing path : %s", b.config.servicePath)
	resp, err := client.Get(b.config.servicePath, false, true)

	if err != nil {
		return err
	}

	for _, n := range resp.Node.Nodes {
		b.Update(&n, resp.Action)
	}

	// Begin the watch after this sync from the next sync
	b.watchIndex = resp.EtcdIndex + 1

	return nil
}

func (b *backends) Next() string {
	b.lock.RLock()
	defer b.lock.RUnlock()

	if len(b.hosts) == 0 {
		return ""
	}

	index := (b.lastIndex + 1) % len(b.hosts)
	b.lastIndex = index

	return b.hosts[index].addr
}
