package spot_consul

type Workload interface {
	Update() error
	GetInstanceLoad() []*InstanceLoad
	GetZoneLoad() map[string]*ZoneLoad
}

type MockWorkload struct {
}

func (wl *MockWorkload) Update() error {
	return nil
}

func (wl *MockWorkload) GetInstanceLoad() []*InstanceLoad {
	var loads []*InstanceLoad
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "1", Ip: "1.1.1.1"})
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "2", Ip: "1.1.1.2"})
	return loads
}

type WorkLoadCpu struct {
	keyUrl string
	loads  []*InstanceLoad
}

func (wl *WorkLoadCpu) Update() error {
	var loads []*InstanceLoad
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "1", Ip: "1.1.1.1"})
	loads = append(loads, &InstanceLoad{Load: 5, InstanceId: "2", Ip: "1.1.1.2"})
	wl.loads = loads
	return nil
}

func (wl *WorkLoadCpu) GetLoad() []*InstanceLoad {
	return wl.loads
}
