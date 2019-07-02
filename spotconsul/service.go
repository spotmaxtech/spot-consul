package spotconsul

import (
	"fmt"
)

// Manage global wide service
type GlobalService struct {
	Services []*Service
}


// Manage a service, every service manage many zones
type Service struct {
	Name   string
	Nodes  map[string]*ServiceNode
	Zones  map[string][]*ServiceNode
	Region string
}

// Better to use a struct to manage node detail
// Data filled when reading consul key
type ServiceNode struct {
	InstanceId string
	Host       string
	Zone       string
}

func GetService(consul *Consul, serviceName string) (* Service, error) {
	entries, err := consul.GetService(serviceName)
	if err != nil {
		return nil, err
	}

	service := &Service {
		Name: serviceName,
	}
	service.Nodes = make(map[string]*ServiceNode)
	service.Zones = make(map[string][]*ServiceNode)

	for _, entry := range entries {
		meta := entry.Service.Meta
		/*
		Meta: {
			balanceFactor: "1900",
			instanceID: "i-09b11dacc2d6e0f2d",
			publicIP: "13.229.182.102",
			zone: "ap-southeast-1b"
		},
		*/
		instanceId, OK := meta["instanceID"]
		if !OK {
			fmt.Println("no instance id meta", entry)
			continue
		}

		zone, OK := meta["zone"]
		if !OK {
			fmt.Println("no zone meta", entry)
			continue
		}

		publicIp, OK := meta["publicIP"]
		if !OK {
			fmt.Println("not public ip", entry)
			publicIp = "unknown"
		}

		node := &ServiceNode{
			InstanceId: instanceId,
			Host: publicIp,
			Zone: zone,
		}
		service.Nodes[instanceId] = node
		service.Zones[zone] = append(service.Zones[zone], node)
	}

	return service, nil
}
