package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	primes         []int
	number         int
	resultExpected []int
}

var tenPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

var inputs = []input{
	{tenPrimes, 30, []int{2, 3, 5}},
	{tenPrimes, 28, []int{2, 2, 7}},
	{tenPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
	{tenPrimes, 30, []int{2, 3, 5}},
	{tenPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
	{tenPrimes, 4, []int{2, 2}},
}

func TestCoveyReverseString(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(FactorNumber(input.primes, input.number), convey.ShouldResemble, input.resultExpected)
			}
		})
	})
}
