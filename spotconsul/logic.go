package spotconsul

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Logic struct {
	Consul        *Consul
	InitialWeight *InitialWeight
	OnlineLab     *OnlineLab
	Service       *Service
	WeightLearner *WeightLearner
	ServiceName   string
	Workload      Workload // interface support other workload

	logicCfg  *LogicConfig
	globalCfg *GlobalConfig
}

func NewLearningLogic(logicCfg *LogicConfig, globalCfg *GlobalConfig) *Logic {
	logic := &Logic{
		Consul:        NewConsul(logicCfg.ConsulAddr),
		OnlineLab:     NewOnlineLab(logicCfg.OnlineLabKey),
		ServiceName:   logicCfg.ServiceName,
		WeightLearner: NewWeightLearner(logicCfg.LearningFactorKey),
		Workload:      NewWorkloadCPU(logicCfg.InstanceLoadKey, logicCfg.ZoneCPUKey),
	}

	logic.logicCfg = logicCfg
	logic.globalCfg = globalCfg
	return logic
}

func (l *Logic) FetchAll() error {
	service, err := GetService(l.Consul, l.ServiceName)
	if err != nil {
		return err
	} else {
		log.Debug("logic service fetched")
		l.Service = service
	}

	if err := l.Workload.Fetch(l.Consul); err != nil {
		return err
	}
	log.Debug("logic workload fetched")

	if err := l.OnlineLab.Fetch(l.Consul); err != nil {
		return err
	}
	log.Debug("logic online lab fetched")

	if err := l.WeightLearner.Fetch(l.Consul); err != nil {
		return nil
	}
	log.Debug("logic weight learner fetched")

	// * 读取默认的权重数据

	return nil
}

func (l *Logic) Learning() error {
	if err := l.WeightLearner.LearningFactors(l.Service, l.Workload, l.OnlineLab); err != nil {
		return err
	}

	if err := l.WeightLearner.LearningCrossRate(l.Workload, l.OnlineLab); err != nil {
		return err
	}

	return nil
}

func (l *Logic) UpdateAll() error {
	if err := l.WeightLearner.Update(l.Consul); err != nil {
		return err
	}
	return nil
}

func (l *Logic) RunOnce() error {
	if err := l.FetchAll(); err != nil {
		return err
	}

	if err := l.Learning(); err != nil {
		return err
	}

	if err := l.UpdateAll(); err != nil {
		return err
	}

	return nil
}

func (l *Logic) RunningLoop(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(l.globalCfg.LoopingTimeS))
	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				log.Infof("service %s logic run once", l.ServiceName)
				if err := l.RunOnce(); err != nil {
					log.Errorf("service %s logic error, %s", l.ServiceName, err.Error())
				}
			case <-ctx.Done():
				log.Infof("service %s logic running loop is canceled", l.ServiceName)
				return
			default:
				log.Debugf("service %s logic running loop sleep", l.ServiceName)
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	wg.Wait()
}
