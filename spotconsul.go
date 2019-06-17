package spot_consul

type SpotConsul struct {
	OnlineLab     *OnlineLab
	Workload      *Workload
	WeightLearner *WeightLearner
	InitialWeight *InitialWeight
}

func (sc *SpotConsul) calculateZoneWeight() (map[string]int64, error) {

	return nil, nil
}

func (sc *SpotConsul) calculateCrossRate() (float32, error) {
	return 0, nil
}

func (sc *SpotConsul) calculateWeightFactors() *WeightFactors {

	factors := &WeightFactors{

	}
	return factors
}

func (sc *SpotConsul) UpdateAll() {
	// 读取consul中的cpu数据，数据是来自cloud watch的
	// 读取online lab中的学习指标数据
	// 读取当前的权重数据
	// * 读取默认的权重数据
	// 按照region组装数据
}

func (sc *SpotConsul) UpdateWeightFactors() {

}

func (sc *SpotConsul) Learning() {

}

func (sc *SpotConsul) Logic() {
	// 获取consul信息
	// 计算每个zone的权值
	// 计算每个region的权值
	// 更新consul信息包括zone和cross
}
