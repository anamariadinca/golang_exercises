package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	list                      []int
	sum, expected1, expected2 int
}

var inputs = []input{
	{[]int{1, 2, 3, 4}, 7, 2, 3},
	{[]int{0, -1, 1}, 0, 1, 2},
	{[]int{0, 1, 1}, 0, -1, -1},
	{[]int{10, 1, 12, 3, 7, 2, 2, 1}, 4, 1, 3},
}

func TestCoveyReverseString(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				index1, index2 := twoThatSum(input.list, input.sum)
				convey.So([2]int{index1, index2}, convey.ShouldEqual, [2]int{input.expected1, input.expected2})
			}
		})
	})
}
