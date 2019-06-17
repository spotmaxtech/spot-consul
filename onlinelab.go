package spot_consul

type CrossZone struct {
	CrossRate    float32
	Enabled      bool
	LearningRate float32
}

type BalanceZone struct {
	Enabled      bool
	LearningRate float32
}

type Lab struct {
	CrossZone   CrossZone
	BalanceZone BalanceZone
}

type OnlineLab struct {
	labUrl string
	Lab    *Lab
}

func NewOnlineLab(url string) *OnlineLab {
	return &OnlineLab{labUrl: url}
}

func (ol *OnlineLab) Update() error {
	lab := &Lab{
		CrossZone:CrossZone{
			CrossRate: 0.05,
			Enabled:true,
			LearningRate: 0.05,
		},
		BalanceZone:BalanceZone{
			Enabled:true,
			LearningRate: 0.05,
		},
	}
	ol.Lab = lab
	return nil
}

func (ol *OnlineLab) GetLab() *Lab {
	return ol.Lab
}
