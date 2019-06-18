package internal

// Manage global wide service
type GlobalService struct {
	Services []*Service
}

type InstanceLoad struct {
	Load       float64
	InstanceId string
	Ip         string
}

// Manage a service, every service manage many zones
type Service struct {
	Name   string
	Nodes  map[string]*ServiceNode
	Zones  []*ServiceZone
	Region string
}

// Better to use a struct to manage node detail
// Data filled when reading consul key
type ServiceNode struct {
	InstanceId string
	Host       string
	Zone       string
}

// Manage a zone, zone manage many nodes
// Zone must contain workload for computing learning metric
type ServiceZone struct {
	Zone     string
	Nodes    []*ServiceNode
}

// All the factors we want
type WeightFactors struct {
	InstanceFactors map[string]float64
	CrossRate       map[string]float64
}

type ZoneLoad struct {
	Zone string
	Load float64
}
