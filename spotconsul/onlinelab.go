package spotconsul

import (
	"encoding/json"
	"errors"
)

// 管理节点间均衡参数
type BalanceZone struct {
	Enabled           bool    `json:"enabled"`
	LearningRate      float64 `json:"learningRate"`
	LearningThreshold float64 `json:"learningThreshold"`
}

// 管理跨区域的均衡参数
type CrossZone struct {
	AlarmThreshold    float64 `json:"alarmThreshold"`
	CrossRate         float64 `json:"crossRate"`
	Enabled           bool    `json:"enabled"`
	LearningRate      float64 `json:"learningRate"`
	LearningThreshold float64 `json:"learningThreshold"`
}

type Lab struct {
	BalanceZone BalanceZone `json:"balanceZone"`
	CrossZone   CrossZone   `json:"crossZone"`
}

type OnlineLab struct {
	key string
	lab *Lab
}

// 需要给个默认的参数
func NewOnlineLab(key string) *OnlineLab {
	ol := &OnlineLab{key: key, lab: nil}
	ol.DefaultLab()
	return ol
}

// 设置默认的lab，方便测试和初期上线，或紧急恢复等等
// 这里一定要设置安全的值
func (ol *OnlineLab) DefaultLab() {
	lab := &Lab{
		BalanceZone: BalanceZone{
			Enabled:           false,
			LearningRate:      0.02,
			LearningThreshold: 0.2,
		},
		CrossZone: CrossZone{
			AlarmThreshold:    70,
			CrossRate:         0.05,
			Enabled:           false,
			LearningRate:      0.02,
			LearningThreshold: 0.2,
		},
	}
	ol.lab = lab
}

// 如果fetch成功了，lab可用不为空
func (ol *OnlineLab) Fetch(consul *Consul) error {
	lab := &Lab{}
	value, err := consul.GetKey(ol.key)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(value, lab); err != nil {
		return err
	}
	ol.lab = lab
	return nil
}

// 使用时防止用到空指针
func (ol *OnlineLab) GetLab() (*Lab, error) {
	if ol.lab == nil {
		return nil, errors.New("lab data is nil, please fetch or set default first")
	}
	return ol.lab, nil
}
