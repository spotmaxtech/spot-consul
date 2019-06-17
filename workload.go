package spot_consul

type Workload interface {
	Update() error
	GetLoad() []*InstanceLoad
}

type MockWorkload struct {
}

func (wl *MockWorkload) Update() error {
	return nil
}

func (wl *MockWorkload) GetLoad() []*InstanceLoad {
	var loads []*InstanceLoad
	loads = append(loads, &InstanceLoad{Factor: 5, InstanceId: "1", Ip: "1.1.1.1"})
	loads = append(loads, &InstanceLoad{Factor: 5, InstanceId: "2", Ip: "1.1.1.2"})
	return loads
}

type WorkLoadCpu struct {
	keyUrl string
	loads  []*InstanceLoad
}

func (wl *WorkLoadCpu) Update() error {
	var loads []*InstanceLoad
	loads = append(loads, &InstanceLoad{Factor: 5, InstanceId: "1", Ip: "1.1.1.1"})
	loads = append(loads, &InstanceLoad{Factor: 5, InstanceId: "2", Ip: "1.1.1.2"})
	wl.loads = loads
	return nil
}

func (wl *WorkLoadCpu) GetLoad() []*InstanceLoad {
	return wl.loads
}
