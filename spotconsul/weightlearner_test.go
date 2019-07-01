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