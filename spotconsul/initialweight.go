package spotconsul

type InitialWeight struct {
	InitialFactors map[string]float64
}

func NewInitialWeight() *InitialWeight {
	initWeight := &InitialWeight{
		InitialFactors: make(map[string]float64),
	}
	return initWeight
}

func (iw *InitialWeight) Fetch(service *Service) {
	initFactors := make(map[string]float64)
	for instanceId, node := range service.Nodes {
		initFactors[instanceId] = node.DefaultFactor
	}
	iw.InitialFactors = initFactors
}

// 计算learning权重和初始化权重，返回确认有风险的实例
// TODO：通过计算方差或人工限定的阈值防止风险
func (iw *InitialWeight) Unhealthy(factors LearningFactors) []string {
	var risks []string
	return risks
}
