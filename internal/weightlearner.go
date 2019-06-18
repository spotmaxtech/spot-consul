package internal

// Weight factor should learn from workload
// 用一个模型来管理权重与学习方法便于升级维护
// 初始化的权重用于控制实时学习权重的偏离风险，明显偏离太多是有可能发生了bug，需要控制一下
type WeightLearner struct {
	KeyUrl          string
	InstanceFactors map[string]float64
	CrossRate       map[string]float64
	InitialWeight   *InitialWeight
}

func (wl *WeightLearner) GetWeightFactors() *WeightFactors {
	factors := &WeightFactors{
		InstanceFactors: wl.InstanceFactors,
		CrossRate:       wl.CrossRate,
	}
	return factors
}

func (wl *WeightLearner) LearningCrossRate(workload Workload, lab *OnlineLab) error {
	zoneLoad := workload.GetZoneLoad()
	var minLoad, maxLoad float64
	var minZone, maxZone string
	for zone, l := range zoneLoad {
		if minLoad > l.Load || minLoad == 0 {
			minLoad = l.Load
			minZone = zone
		}
		if maxLoad < l.Load || maxLoad == 0 {
			maxLoad = l.Load
			maxZone = zone
		}
	}

	if minZone != maxZone && maxLoad-minLoad > lab.Lab.CrossZone.LearningThreshold {
		wl.CrossRate[maxZone] -= wl.CrossRate[maxZone] * lab.Lab.CrossZone.LearningRate
		wl.CrossRate[minZone] += wl.CrossRate[minZone] * lab.Lab.CrossZone.LearningRate
	} else {
		for zone := range zoneLoad {
			wl.CrossRate[zone] -= wl.CrossRate[zone] * lab.Lab.CrossZone.LearningRate
		}
	}
	return nil
}

func (wl *WeightLearner) LearningFactors(service *Service, workload Workload, lab *OnlineLab) error {
	zoneLoad := workload.GetZoneLoad()
	instanceLoad := workload.GetInstanceLoad()

	for _, l := range instanceLoad {
		node := service.Nodes[l.InstanceId]
		zone := node.Zone
		if l.Load > zoneLoad[zone].Load*(1+lab.Lab.BalanceZone.LearningThreshold) {
			wl.InstanceFactors[l.InstanceId] -= wl.InstanceFactors[l.InstanceId] * lab.Lab.BalanceZone.LearningRate
		} else if l.Load < zoneLoad[zone].Load*(1-lab.Lab.BalanceZone.LearningThreshold) {
			wl.InstanceFactors[l.InstanceId] += wl.InstanceFactors[l.InstanceId] * lab.Lab.BalanceZone.LearningRate
		}
	}
	return nil
}

func (wl *WeightLearner) Fetch() error {
	// TODO: fetch consul data
	return nil
}

func (wl *WeightLearner) Update() error {
	// TODO: update back to consul
	return nil

}
