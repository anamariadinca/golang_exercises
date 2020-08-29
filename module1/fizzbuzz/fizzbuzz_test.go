package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	number int
	want   string
}

var inputs = []input{
	{1, "1\n"},
	{2, "1, 2\n"},
	{3, "1, 2, Fizz\n"},
	{4, "1, 2, Fizz, 4\n"},
	{5, "1, 2, Fizz, 4, Buzz\n"},
	{6, "1, 2, Fizz, 4, Buzz, Fizz\n"},
	{7, "1, 2, Fizz, 4, Buzz, Fizz, 7\n"},
	{8, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8\n"},
	{9, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz\n"},
	{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz\n"},
	{11, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11\n"},
	{12, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz\n"},
	{13, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13\n"},
	{14, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14\n"},
	{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz\n"},
	{16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16\n"},
	{17, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17\n"},
	{18, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17, Fizz\n"},
	{19, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17, Fizz, 19\n"},
	{20, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17, Fizz, 19, Buzz\n"},
}

func TestConveyFizzBuzz(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(fizzBuzz(input.number), convey.ShouldEqual, input.want)
			}
		})
	})
}
