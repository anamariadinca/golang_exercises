package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	list     []int
	sum      int
	expected bool
}

var inputs = []input{
	{[]int{1, 2, 3, 4}, 7, true},
	{[]int{0, -1, 1}, 0, true},
	{[]int{0, 1, 1}, 0, false},
	{[]int{10, 1, 12, 3, 7, 2, 2, 1}, 4, true},
}

func TestCoveyReverseString(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				index1, index2 := twoThatSum(input.list, input.sum)
				result := decideResult(index1, index2)
				convey.So(result, convey.ShouldEqual, input.expected)
			}
		})
	})
}

func decideResult(ind1 int, ind2 int) bool {
	var result bool
	if ind1 != -1 && ind2 != -1 {
		result = true
	} else {
		result = false
	}
	return result
}
