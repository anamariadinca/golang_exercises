package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	pos, expected int
}

var inputs = []input{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{14, 377},
	{22, 17711},
	{25, 75025},
	// This test case may be much slower depending on your
	// solution. We will look at how to speed it up in a future
	// module. Feel free to comment it out if it is too slow.
	{41, 165580141},
}

func TestCoveyFibonacci(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(getFibonacciForGivenPosition(input.pos), convey.ShouldResemble, input.expected)
			}
		})
	})
}
