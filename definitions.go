package spot_consul

// Manage global wide service
type Global struct {
	Services []*Service
}

type InstanceLoad struct {
	Factor     int32
	InstanceId string
	Ip         string
}

// Manage a service, every service manage many zones
type Service struct {
	Name   string
	Region string
	Zones  []*ServiceZone
}

// Better to use a struct to manage node detail
// Data filled when reading consul key
type ServiceNode struct {
	Factor     int64
	InstanceId string
	Host       string
	Workload   int64
	Zone       string
}

// Manage a zone, zone manage many nodes
// Zone must contain workload for computing learning metric
type ServiceZone struct {
	Zone     string
	Workload int64
	Nodes    []*ServiceNode
}

// All the factors we want
type WeightFactors struct {
	InstanceFactors map[string]int64
	CrossRate       float32
}
