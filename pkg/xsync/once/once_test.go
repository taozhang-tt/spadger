package once

import (
	"errors"
	"sync"
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestOnce(t *testing.T) {
	Convey("TestOnce", t, func() {
		var once Once

		counter := 0

		f := func() error {
			counter++
			return nil
		}

		wg := sync.WaitGroup{}
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				once.Do(f)
				wg.Done()
			}()
		}
		wg.Wait()
		So(counter, ShouldEqual, 1)
	})

	Convey("TestError", t, func() {
		var once Once

		counter := 0

		f := func() error {
			counter++
			if counter > 5 {
				return nil
			}
			return errors.New("error")
		}

		for i := 1; i < 10; i++ {
			err := once.Do(f)
			if i <= 5 {
				So(err, ShouldNotBeNil)
				So(once.Done(), ShouldBeFalse)
				So(counter, ShouldEqual, i)
			} else {
				So(err, ShouldBeNil)
				So(once.Done(), ShouldBeTrue)
				So(counter, ShouldEqual, 6)
			}
		}
	})
}
