package coll

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColl(t *testing.T) {
	Convey("Coll Test", t, func() {
		Convey("Map2UrlEncodedString test", func() {
			m := MapOf("q[id]", "3", "q[name::like]", "a2d%")
			s := Map2UrlEncodedString(m)
			m2 := UrlEncodedString2Map(s)
			So(Map2SortString(m), ShouldEqual, Map2SortString(m2))
		})
	})
}

func TestEqJoinedSplitter(t *testing.T) {
	k, v := EqJoinedKVSplitter("a3")
	fmt.Println(k, v)
}
