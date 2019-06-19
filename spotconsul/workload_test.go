package spotconsul

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewWorkloadCPU(t *testing.T) {
	Convey("test new workload", t, func() {
		consul := NewConsul(TestConsulAddress)
		workload := NewWorkloadCPU(TestInstanceFactorKey, TestZoneCPUKey)
		err := workload.Fetch(consul)
		So(err, ShouldBeNil)
		So(workload.zoneLoadKey, ShouldEqual,TestZoneCPUKey)
		t.Log(Prettify(workload.zoneLoads))
		t.Log(Prettify(workload.instanceLoads))
	})
}
