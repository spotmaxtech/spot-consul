package internal

type SpotConsul struct {
	Service       *Service
	OnlineLab     *OnlineLab
	Workload      *Workload
	InitialWeight *InitialWeight
	WeightLearner *WeightLearner
}

func (sc *SpotConsul) FetchAll() error {
	// 获取当前的全局服务节点
	// 读取consul中的cpu数据，数据是来自cloud watch的
	// 读取online lab中的学习指标数据
	// 读取当前的权重数据
	// * 读取默认的权重数据
	// 按照region组装数据
	return nil
}

func (sc *SpotConsul) Learning() error {
	if err := sc.WeightLearner.LearningFactors(sc.Service, *sc.Workload, sc.OnlineLab); err != nil {
		return err
	}

	if err := sc.WeightLearner.LearningCrossRate(*sc.Workload, sc.OnlineLab); err != nil {
		return err
	}

	return nil
}

func (sc *SpotConsul) UpdateAll() error {
	if err := sc.WeightLearner.Update(); err != nil {
		return err
	}
	return nil
}

func (sc *SpotConsul) Logic() error {
	if err := sc.FetchAll(); err != nil {
		return err
	}
	if err := sc.Learning(); err != nil {
		return err
	}

	if err := sc.UpdateAll(); err != nil {
		return err
	}

	return nil
}
