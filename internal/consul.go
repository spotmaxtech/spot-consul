package internal

import "github.com/hashicorp/consul/api"

type Consul struct {
	address string
	health  *api.Health
	kv      *api.KV
}

// 创建一个链接，失败的话会panic
func NewConsul(address string) *Consul {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		panic("connect consul failed, err:" + err.Error())
	}

	consul := &Consul{
		address: address,
		health:  client.Health(),
		kv:      client.KV(),
	}
	return consul
}

func (c *Consul) GetKey(key string) ([]byte, error) {
	pair, _, err := c.kv.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return pair.Value, nil
}

// 包装一下，不需要meta，只要健康的节点
func (c *Consul) GetService(service string) ([]*api.ServiceEntry, error) {
	entry, _, err := c.health.Service(service, "", true, nil)
	return entry, err
}

func (c *Consul) PutKey(key string, value string) error {
	pair := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := c.kv.Put(pair, nil)
	return err
}
