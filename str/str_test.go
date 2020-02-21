package str

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	Convey("Join Test", t, func() {
		So(Join(","), ShouldEqual, "")
		So(Join(",", "abc"), ShouldEqual, "abc")
		So(Join(",", "abc", "def"), ShouldEqual, "abc,def")
		So(Join(",", "abc", "def", "gh"), ShouldEqual, "abc,def,gh")
	})
}

func TestRandStr(t *testing.T) {
	s := RandStrWithCharset(10, Letters)
	assert.Equal(t, len(s), 10)
}
