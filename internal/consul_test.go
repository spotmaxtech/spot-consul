package internal

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestConsul_KV(t *testing.T) {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "13.251.183.9:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "spotmax/test", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("spotmax/test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

func TestConsul_Health(t *testing.T) {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "13.251.183.9:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	health := client.Health()
	entry, meta, err := health.Service("rs", "", true, nil)
	t.Log(entry, meta)
}

func TestString(t *testing.T) {
	s := String("a")
	t.Log(s)
}