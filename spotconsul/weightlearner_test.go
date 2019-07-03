package spotconsul

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
	"time"
)

func TestWeightLearner_Fetch(t *testing.T) {
	Convey("Test learner fetch", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := NewWeightLearner(TestLearningFactorsKey)
		err := learner.Fetch(consul)
		So(err, ShouldBeNil)
		t.Log(learner.Factors)
	})
}

func TestWeightLearner_Update(t *testing.T) {
	Convey("Test learner update", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := MockWeightLearner(TestLearningFactorsKey)
		err := learner.Update(consul)
		So(err, ShouldBeNil)
	})
}

func TestWeightLearner_LearningOnce(t *testing.T) {
	Convey("Test learning", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := NewWeightLearner(TestLearningFactorsKey)
		err := learner.Fetch(consul)
		So(err, ShouldBeNil)
		t.Log(learner.Factors)

		ol := NewOnlineLab(TestOnlineLabKey)
		So(ol.key, ShouldEqual, TestOnlineLabKey)
		t.Log(ol.lab)

		wl := &MockWorkload{}
		instLoad := wl.GetInstanceLoad()
		t.Log(instLoad)
		zoneLoad := wl.GetZoneLoad()
		t.Log(zoneLoad)

		service := &Service{
			Name: "rs",
		}
		nodes := make(map[string]*ServiceNode)
		nodes["i-1"] = &ServiceNode{
			DefaultFactor: 1000,
			InstanceId:    "i-1",
			Host:          "1.1.1.1",
			Zone:          "us-west-2a",
		}
		nodes["i-2"] = &ServiceNode{
			DefaultFactor: 800,
			InstanceId:    "i-2",
			Host:          "2.2.2.2",
			Zone:          "us-west-2b",
		}
		service.Nodes = nodes

		err = learner.LearningFactors(service, wl, ol)
		So(err, ShouldBeNil)
		t.Log(Prettify(learner.Factors))

		err = learner.LearningCrossRate(wl, ol)
		So(err, ShouldBeNil)
		t.Log(Prettify(learner.Factors))
	})
}

func TestWeightLearner_InitialLearning(t *testing.T) {
	Convey("Test learning", t, func() {
		logrus.SetLevel(logrus.DebugLevel)
		consul := NewConsul(TestConsulAddress)
		TempTestKey := "spotmax-test/learning_factor2.json"
		learner := NewWeightLearner(TempTestKey)
		err := learner.Fetch(consul)
		So(err, ShouldBeNil)
		t.Log(learner.Factors)

		ol := NewOnlineLab(TestOnlineLabKey)
		So(ol.key, ShouldEqual, TestOnlineLabKey)
		t.Logf("%#v", ol.lab)

		wl := NewWorkloadCPU(TestInstanceFactorKey, TestZoneCPUKey)
		err = wl.Fetch(consul)
		So(err, ShouldBeNil)
		instLoad := wl.GetInstanceLoad()
		t.Logf("%#v", instLoad)
		zoneLoad := wl.GetZoneLoad()
		t.Logf("%#v", zoneLoad)

		service := &Service{
			Name: "rs",
		}
		nodes := make(map[string]*ServiceNode)
		nodes["i-1"] = &ServiceNode{
			DefaultFactor: 1000,
			InstanceId:    "i-1",
			Host:          "1.1.1.1",
			Zone:          "us-west-2a",
		}
		nodes["i-2"] = &ServiceNode{
			DefaultFactor: 800,
			InstanceId:    "i-2",
			Host:          "2.2.2.2",
			Zone:          "us-west-2b",
		}
		service.Nodes = nodes

		err = learner.LearningFactors(service, wl, ol)
		So(err, ShouldBeNil)
		t.Log(Prettify(learner.Factors))

		err = learner.LearningCrossRate(wl, ol)
		So(err, ShouldBeNil)
		t.Log(Prettify(learner.Factors))

		err = learner.Update(consul)
		So(err, ShouldBeNil)

		_, err = consul.kv.Delete(TempTestKey, nil)
		So(err, ShouldBeNil)
	})
}

func TestWeightLearner_LearningLoop(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(5 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("pat")
			case <-timer.C:
				fmt.Println("timeout")
				ticker.Stop()
				wg.Done()
				return
			}
		}

	}()

	wg.Wait()
}
