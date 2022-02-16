package xtime

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestTimezone(t *testing.T) {
	Convey("Test Timezone", t, func() {
		loc, err := GetLocation("Asia/Shanghai")
		So(err, ShouldBeNil)
		So(loc.String(), ShouldEqual, "Asia/Shanghai")

		loc, err = GetLocation("Asia/Shanghai_NotExist")
		So(loc, ShouldBeNil)
		So(err.Error(), ShouldEqual, "unknown time zone Asia/Shanghai_NotExist")

		loc, err = GetLocation("Asia/Shanghai_NotExist")
		So(loc, ShouldBeNil)
		So(err, ShouldResemble, errCacheLocation)
	})
}
