package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	list           []int
	expectedResult int
}

var listInputs = []input{
	{[]int{2, 3, 5}, 10},
	{[]int{10, 22, 4}, 36},
	{[]int{1, 0, -3}, -2},
	{nil, 0},
	{[]int{}, 0},
}

func TestConveySumList(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range listInputs {
				convey.So(ComputeSumOfList(input.list), convey.ShouldEqual, input.expectedResult)
			}
		})
	})
}

func TestConveySumListRecursive(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range listInputs {
				convey.So(ComputeSumRecursive(input.list), convey.ShouldEqual, input.expectedResult)
			}
		})
	})
}
