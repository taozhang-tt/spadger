package xheap

import (
	"math/rand"
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestPriorityQueue(t *testing.T) {
	Convey("PriorityQueue", t, func() {
		max := 10000
		for turn := 0; turn < 2; turn++ {
			min := turn == 0
			pq := NewPriorityQueue(32, min)

			So(pq.Len(), ShouldEqual, 0)
			So(pq.Cap(), ShouldEqual, 32)
			So(pq.Empty(), ShouldBeTrue)

			for i := 0; i < 64; i++ {
				pq.Push(i, rand.Intn(max))
			}

			So(pq.Len(), ShouldEqual, 64)
			So(pq.Cap(), ShouldEqual, 64)
			So(pq.Empty(), ShouldBeFalse)

			pre := 0
			if min {
				pre = -1
			} else {
				pre = max
			}
			for i := 0; i < 64; i++ {
				v1, s1 := pq.Top()
				v2, s2 := pq.Pop()
				So(pq.Len(), ShouldEqual, 63-i)
				So(v1, ShouldEqual, v2)
				So(s1, ShouldEqual, s2)
				if min {
					So(pre, ShouldBeLessThanOrEqualTo, s1)
				} else {
					So(pre, ShouldBeGreaterThanOrEqualTo, s1)
				}
				pre = s1
			}
		}
	})

	Convey("PriorityQueue.New", t, func() {
		pq1 := NewPriorityQueue(32)
		So(pq1.priorityQueue.min, ShouldBeTrue)
		pq2 := NewPriorityQueue(31, true)
		So(pq2.priorityQueue.min, ShouldBeTrue)
		pq3 := NewPriorityQueue(32, false)
		So(pq3.priorityQueue.min, ShouldBeFalse)
	})
}

func TestPriorityQueue64(t *testing.T) {
	Convey("PriorityQueue64", t, func() {
		max := 10000
		for turn := 0; turn < 2; turn++ {
			min := turn == 0
			pq := NewPriorityQueue64(32, min)

			So(pq.Len(), ShouldEqual, 0)
			So(pq.Cap(), ShouldEqual, 32)
			So(pq.Empty(), ShouldBeTrue)

			for i := 0; i < 64; i++ {
				pq.Push(i, int64(rand.Intn(max)))
			}

			So(pq.Len(), ShouldEqual, 64)
			So(pq.Cap(), ShouldEqual, 64)
			So(pq.Empty(), ShouldBeFalse)

			var pre int64
			if min {
				pre = -1
			} else {
				pre = int64(max)
			}
			for i := 0; i < 64; i++ {
				v1, s1 := pq.Top()
				v2, s2 := pq.Pop()
				So(pq.Len(), ShouldEqual, 63-i)
				So(v1, ShouldEqual, v2)
				So(s1, ShouldEqual, s2)
				if min {
					So(pre, ShouldBeLessThanOrEqualTo, s1)
				} else {
					So(pre, ShouldBeGreaterThanOrEqualTo, s1)
				}
				pre = s1
			}
		}
	})

	Convey("PriorityQueue64.New", t, func() {
		pq1 := NewPriorityQueue64(32)
		So(pq1.priorityQueue64.min, ShouldBeTrue)
		pq2 := NewPriorityQueue64(32, true)
		So(pq2.priorityQueue64.min, ShouldBeTrue)
		pq3 := NewPriorityQueue64(32, false)
		So(pq3.priorityQueue64.min, ShouldBeFalse)
	})
}
