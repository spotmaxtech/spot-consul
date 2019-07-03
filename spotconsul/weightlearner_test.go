package spotconsul

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWeightLearner_Update(t *testing.T) {
	Convey("Test learner update", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := MockWeightLearner(TestLearningFactorsKey)
		err := learner.Update(consul)
		So(err, ShouldBeNil)
	})
}

func TestWeightLearner_Fetch(t *testing.T) {
	Convey("Test learner fetch", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := NewWeightLearner(TestLearningFactorsKey)
		err := learner.Fetch(consul)
		So(err, ShouldBeNil)
		t.Log(learner.Factors)
	})
}

func TestWeightLearner_LearningFactors(t *testing.T) {
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
			InstanceId:"i-1",
			Host:"1.1.1.1",
			Zone:"us-west-2a",
		}
		nodes["i-2"] = &ServiceNode{
			InstanceId:"i-2",
			Host:"2.2.2.2",
			Zone:"us-west-2b",
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


func TestWeightLearner_Online(t *testing.T) {
	Convey("Test learning", t, func() {
		consul := NewConsul(TestConsulAddress)
		learner := NewWeightLearner("spotmax-test/learning_factor2.json")
		err := learner.Fetch(consul)
		So(err, ShouldBeNil)
		// learner.Factors.InstanceFactors["i-1"] = 100
		// learner.Factors.CrossRate["z-1"] = 100
		t.Log(learner.Factors)

		// ol := NewOnlineLab(TestOnlineLabKey)
		// So(ol.key, ShouldEqual, TestOnlineLabKey)
		// t.Log(ol.lab)
		//
		// wl := NewWorkloadCPU(TestInstanceFactorKey, TestZoneCPUKey)
		// err = wl.Fetch(consul)
		// So(err, ShouldBeNil)
		// instLoad := wl.GetInstanceLoad()
		// t.Log(instLoad)
		// zoneLoad := wl.GetZoneLoad()
		// t.Log(zoneLoad)
		//
		// service := &Service{
		// 	Name: "rs",
		// }
		// nodes := make(map[string]*ServiceNode)
		// nodes["i-1"] = &ServiceNode{
		// 	InstanceId:"i-1",
		// 	Host:"1.1.1.1",
		// 	Zone:"us-west-2a",
		// }
		// nodes["i-2"] = &ServiceNode{
		// 	InstanceId:"i-2",
		// 	Host:"2.2.2.2",
		// 	Zone:"us-west-2b",
		// }
		// service.Nodes = nodes
		//
		// err = learner.LearningFactors(service, wl, ol)
		// So(err, ShouldBeNil)
		// t.Log(Prettify(learner.Factors))
		//
		// err = learner.LearningCrossRate(wl, ol)
		// So(err, ShouldBeNil)
		// t.Log(Prettify(learner.Factors))
	})
}