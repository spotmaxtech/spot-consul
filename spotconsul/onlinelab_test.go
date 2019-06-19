package spotconsul

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewOnlineLab(t *testing.T) {
	Convey("test new", t, func() {
		ol := NewOnlineLab(TestOnlineLabKey)
		So(ol.key, ShouldEqual, TestOnlineLabKey)
		So(ol.lab.BalanceZone.Enabled, ShouldEqual, false)
	})
}

func TestOnlineLab_Fetch(t *testing.T) {
	Convey("test onlinelab", t, func() {
		consul := NewConsul(TestConsulAddress)
		ol := NewOnlineLab(TestOnlineLabKey)

		// set one only for test
		lab1 := Lab{
			BalanceZone: BalanceZone{
				Enabled:           true,
				LearningRate:      0.1,
				LearningThreshold: 0.2,
			},
			CrossZone: CrossZone{
				AlarmThreshold:    60,
				CrossRate:         0.05,
				Enabled:           false,
				LearningRate:      0.1,
				LearningThreshold: 0.2,
			},
		}
		bytes, err := json.MarshalIndent(lab1, "", "    ")
		So(err, ShouldBeNil)
		err = consul.PutKey(TestOnlineLabKey, bytes)
		So(err, ShouldBeNil)

		// get
		err = ol.Fetch(consul)
		So(err, ShouldBeNil)
		So(ol.lab.BalanceZone.Enabled, ShouldEqual, true)
		So(ol.lab.BalanceZone.LearningRate, ShouldEqual, 0.1)
		So(ol.lab.BalanceZone.LearningThreshold, ShouldEqual, 0.2)

		// delete
		_, err = consul.kv.Delete(TestOnlineLabKey, nil)
		So(err, ShouldBeNil)
	})
}
