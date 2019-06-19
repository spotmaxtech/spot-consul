package spotconsul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const TestConsulAddress  = "13.251.183.9:8500"
const TestOnlineLabKey = "spotmax-test/onlinelab.json"

func TestConsul_KV(t *testing.T) {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = TestConsulAddress
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "spotmax-test/test", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("spotmax-test/test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

func TestConsul_Health(t *testing.T) {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = TestConsulAddress
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	health := client.Health()
	entry, meta, err := health.Service("rs", "", true, nil)
	t.Log(Prettify(entry), Prettify(meta))
}

func TestString(t *testing.T) {
	s := String("a")
	t.Log(s)
}

func TestConsul_GetKey(t *testing.T) {
	Convey("test get key", t, func() {
		consul := NewConsul(TestConsulAddress)
		key := "spotmax-test/foo"
		err := consul.PutKey(key, []byte("bar"))
		So(err, ShouldBeNil)
		value, err := consul.GetKey(key)
		So(err, ShouldBeNil)
		So(string(value), ShouldEqual, "bar")
		_, err = consul.kv.Delete(key, nil)
		So(err, ShouldBeNil)
	})
}

func TestConsul_GetService(t *testing.T) {
	Convey("test get service", t, func() {
		consul := NewConsul(TestConsulAddress)
		entry, err := consul.GetService("rs")
		So(err, ShouldBeNil)
		So(len(entry), ShouldBeGreaterThan, 0)
		entry, err = consul.GetService("as")
		So(err, ShouldBeNil)
		So(len(entry), ShouldBeGreaterThan, 0)
		t.Log(Prettify(entry))
	})
}
