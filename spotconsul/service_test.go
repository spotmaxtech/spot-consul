package spotconsul

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetService(t *testing.T) {
	Convey("Test get service", t, func() {
		consul := NewConsul(TestConsulAddress)
		rsService, err := GetService(consul, "rs")
		So(err, ShouldBeNil)
		t.Log(Prettify(rsService))

		asService, err := GetService(consul, "as")
		So(err, ShouldBeNil)
		t.Log(Prettify(asService))
	})
}
