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

		wl := MockWorkload{}
		instLoad := wl.GetInstanceLoad()
		t.Log(instLoad)
		zoneLoad := wl.GetZoneLoad()
		t.Log(zoneLoad)

		// learner.LearningFactors()
	})
}
