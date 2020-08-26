package module1

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type input struct {
	word         string
	reversedWord string
}

var inputs = []input{
	{"abc", "cba"},
}

func TestCoveyReverseString(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range inputs {
				convey.So(reverseString(input.word), convey.ShouldEqual, input.reversedWord)
			}
		})
	})
}
