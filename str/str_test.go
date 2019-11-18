package str

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJoin(t *testing.T) {
	Convey("Join Test", t, func() {
		So(Join(","), ShouldEqual, "")
		So(Join(",", "abc"), ShouldEqual, "abc")
		So(Join(",", "abc", "def"), ShouldEqual, "abc,def")
		So(Join(",", "abc", "def", "gh"), ShouldEqual, "abc,def,gh")
	})
}
