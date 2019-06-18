package internal

type CrossZone struct {
	Enabled           bool
	CrossRate         float64
	AlarmThreshold    float64
	LearningThreshold float64
	LearningRate      float64
}

type BalanceZone struct {
	Enabled           bool
	LearningThreshold float64
	LearningRate      float64
}

type Lab struct {
	CrossZone   CrossZone
	BalanceZone BalanceZone
}

type OnlineLab struct {
	KeyUrl string
	Lab    *Lab
}

func NewOnlineLab(url string) *OnlineLab {
	return &OnlineLab{KeyUrl: url}
}

func (ol *OnlineLab) Update() error {
	lab := &Lab{
		CrossZone: CrossZone{
			CrossRate:    0.05,
			Enabled:      true,
			LearningRate: 0.05,
		},
		BalanceZone: BalanceZone{
			Enabled:      true,
			LearningRate: 0.05,
		},
	}
	ol.Lab = lab
	return nil
}

func (ol *OnlineLab) GetLab() *Lab {
	return ol.Lab
}
