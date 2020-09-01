package module2

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	list, expected []int
}

var inputs = []input{
	{[]int{0, 1}, []int{0, 1}},
	{[]int{1, 0, 2}, []int{0, 1, 2}},
}

func TestCoveyBubbleSort(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(BubbleSort(input.list), convey.ShouldResemble, input.expected)
			}
		})
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, size := range []int{100, 200, 300, 800, 1600} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				b.StartTimer()
				BubbleSort(list)
			}
		})
	}
}
