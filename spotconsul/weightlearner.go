package spotconsul

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type LearningFactors struct {
	InstanceFactors map[string]float64	`json:"instanceFactors"`
	CrossRate       map[string]float64	`json:"crossRate"`
}

// Weight factor should learn from workload
// 用一个模型来管理权重与学习方法便于升级维护
// 初始化的权重用于控制实时学习权重的偏离风险，明显偏离太多是有可能发生了bug，需要控制一下
type WeightLearner struct {
	Key           string
	Factors       *LearningFactors
	InitialWeight *InitialWeight
}

func NewWeightLearner(key string) *WeightLearner {
	learner := &WeightLearner{
		Key: key,
	}

	return learner
}

// 方便做前期测试使用，设置默认的学习数据
func MockWeightLearner(key string) *WeightLearner {
	learner := &WeightLearner{
		Key: key,
	}

	factors := &LearningFactors{}
	factors.InstanceFactors = make(map[string]float64)
	factors.InstanceFactors["i-1"] = 1500
	factors.InstanceFactors["i-2"] = 1200

	factors.CrossRate = make(map[string]float64)
	factors.CrossRate["us-west-2a"] = 0.1
	factors.CrossRate["us-west-2b"] = 0.2

	learner.Factors = factors
	return learner
}

func (wl *WeightLearner) GetLearningFactors() *LearningFactors {
	return wl.Factors
}

func (wl *WeightLearner) LearningCrossRate(workload Workload, ol *OnlineLab) error {
	zoneLoad := workload.GetZoneLoad()
	var minLoad, maxLoad float64
	var minZone, maxZone string
	for zone, load := range zoneLoad {
		if minLoad > load || minLoad == 0 {
			minLoad = load
			minZone = zone
		}
		if maxLoad < load || maxLoad == 0 {
			maxLoad = load
			maxZone = zone
		}
	}

	if minZone != maxZone && maxLoad-minLoad > ol.lab.CrossZone.LearningThreshold {
		wl.Factors.CrossRate[maxZone] -= wl.Factors.CrossRate[maxZone] * ol.lab.CrossZone.LearningRate
		wl.Factors.CrossRate[minZone] += wl.Factors.CrossRate[minZone] * ol.lab.CrossZone.LearningRate
	} else {
		for zone := range zoneLoad {
			wl.Factors.CrossRate[zone] -= wl.Factors.CrossRate[zone] * ol.lab.CrossZone.LearningRate
		}
	}
	return nil
}

func (wl *WeightLearner) LearningFactors(service *Service, workload Workload, ol *OnlineLab) error {
	zoneLoad := workload.GetZoneLoad()
	instanceLoad := workload.GetInstanceLoad()

	for _, l := range instanceLoad {
		node, OK := service.Nodes[l.InstanceId]
		if !OK {
			fmt.Println("can not find proper node", l.InstanceId)
			continue
		}
		zone := node.Zone
		if l.Load > zoneLoad[zone]*(1+ol.lab.BalanceZone.LearningThreshold) {
			wl.Factors.InstanceFactors[l.InstanceId] -= wl.Factors.InstanceFactors[l.InstanceId] * ol.lab.BalanceZone.LearningRate
		} else if l.Load < zoneLoad[zone]*(1-ol.lab.BalanceZone.LearningThreshold) {
			wl.Factors.InstanceFactors[l.InstanceId] += wl.Factors.InstanceFactors[l.InstanceId] * ol.lab.BalanceZone.LearningRate
			fmt.Printf("increase instance %s by %g \n", l.InstanceId, ol.lab.BalanceZone.LearningRate)
		}
	}
	return nil
}

func (wl *WeightLearner) Fetch(consul *Consul) error {
	factorsValue, err := consul.GetKey(wl.Key)
	if err != nil {
		return errors.Errorf("learning factors get failed, %s", err.Error())
	}

	var factors LearningFactors

	if err := json.Unmarshal(factorsValue, &factors); err != nil {
		return errors.Errorf("unmarshal failed for learning factors, %s", err.Error())
	}

	wl.Factors = &factors
	return nil
}

func (wl *WeightLearner) Update(consul *Consul) error {
	factorsValue, err := json.MarshalIndent(wl.Factors, "", "    ")
	if err != nil {
		return errors.Errorf("marshal failed for instance factors %s", err.Error())
	}

	if err := consul.PutKey(wl.Key, factorsValue); err != nil {
		return err
	}

	return nil
}
