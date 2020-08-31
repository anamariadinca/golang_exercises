package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	first, second, cod int
}

var inputs = []input{
	{10, 5, 5},
	{25, 5, 5},
	{30, 15, 15},
	{30, 9, 3},
	{100, 9, 1},
}

func TestCoveyGCDFactoring(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(gcdUsingFactoring(input.first, input.second), convey.ShouldResemble, input.cod)
			}
		})
	})
}

func TestCoveyGCDEuclid(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(gcdUsingEuclidean(input.first, input.second), convey.ShouldResemble, input.cod)
			}
		})
	})
}
