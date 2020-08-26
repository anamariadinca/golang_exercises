package module1

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

type input struct {
	testNum []int
	searchNumber int
	expected bool
}

var inputs = []input{
	{[]int{2,3,5,10,22,13,4},13,true},
	{[]int{2,3,5,10,22,13,4},15,false},
	{[]int{1,0,-3,2,3,5,10,22,13,4},-3,true},
	{nil, 5, false,},
	{[]int{}, 10, false,},
}

func TestConveyTestFindNumber (t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(NumInList(input.testNum, input.searchNumber), convey.ShouldEqual, input.expected)
			}
		})
	})
}