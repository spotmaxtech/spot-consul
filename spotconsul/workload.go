package spotconsul

import "encoding/json"

type InstanceLoad struct {
	Load       float64
	InstanceId string
	PublicIp   string
}

type CDInstanceFactor struct {
	CPUUtilization float64 `json:"CPUUtilization"`
	InstanceId     string  `json:"instanceid"`
	PublicIp       string  `json:"public_ip"`
	Zone           string  `json:"zone"`
}

type CDFactors struct {
	Updated int64               `json:"updated"`
	Data    []*CDInstanceFactor `json:"data"`
}

type CDZoneCPU struct {
	Updated int64         `json:"updated"`
	Data    []interface{} `json:"data"`
}

type DataInstanceFactor struct {
}

type Workload interface {
	GetInstanceLoad() []*InstanceLoad
	GetZoneLoad() map[string]float64
}

type WorkloadCPU struct {
	instanceLoadKey string
	instanceLoads   []*InstanceLoad
	zoneLoadKey     string
	zoneLoads       map[string]float64
}

func NewWorkloadCPU(instanceLoadKey string, zoneLoadKey string) *WorkloadCPU {
	workload := &WorkloadCPU{
		instanceLoadKey: instanceLoadKey,
		zoneLoadKey:     zoneLoadKey,
	}
	return workload
}

func (wl *WorkloadCPU) Fetch(consul *Consul) error {
	// instance load
	value, err := consul.GetKey(wl.instanceLoadKey)
	if err != nil {
		return err
	}
	factors := CDFactors{}
	if err := json.Unmarshal(value, &factors); err != nil {
		return err
	}

	var loads []*InstanceLoad
	for _, factor := range factors.Data {
		loads = append(loads, &InstanceLoad{
			Load:       factor.CPUUtilization,
			InstanceId: factor.InstanceId,
			PublicIp:   factor.PublicIp,
		})
	}

	wl.instanceLoads = loads

	// zone load
	value, err = consul.GetKey(wl.zoneLoadKey)
	if err != nil {
		return err
	}
	zoneCPU := CDZoneCPU{}
	if err := json.Unmarshal(value, &zoneCPU); err != nil {
		return err
	}

	zoneLoads := make(map[string]float64)
	for _, zone := range zoneCPU.Data {
		for k, v := range zone.(map[string]interface{}) {
			zoneLoads[k] = v.(float64)
		}
	}

	wl.zoneLoads = zoneLoads

	return nil
}

func (wl *WorkloadCPU) GetInstanceLoad() []*InstanceLoad {
	return wl.instanceLoads
}

func (wl *WorkloadCPU) GetZoneLoad() map[string]float64 {
	return wl.zoneLoads
}

// Mock一个负载对象用于快速开发测试
type MockWorkload struct {
}

func (wl *MockWorkload) GetInstanceLoad() []*InstanceLoad {
	var loads []*InstanceLoad
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "i-1", PublicIp: "1.1.1.1"})
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "i-2", PublicIp: "1.1.1.2"})
	return loads
}

func (wl *MockWorkload) GetZoneLoad() map[string]float64 {
	load := make(map[string]float64)
	load["us-west-2a"] = 60
	load["us-west-2b"] = 30
	return load
}
