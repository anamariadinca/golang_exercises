package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	// "fmt"
)

type numberValidElement struct {
	inputedNumber  int
	expected bool
}

type compareNumbersElement struct {
	guessedNumber int
	numberGenerated int
	noOfTries int
	expectedResult bool
	isGuessHigher bool
	anyTriesLeft bool
}

var numberValidationTest = []numberValidElement{
	{15, true},
	{-3, false},
	{101, false},
}

var compareNumbersTest = []compareNumbersElement{
	{15, 34, 9, true, false, false},
	{-3, 56, 4, false, false, true},
	{101, 22, 2, false, true, true},
	{24, 23, 5, true, true, true},
}

func TestConveyNumberValid(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range numberValidationTest {
				convey.So(isNumberValid(input.inputedNumber), convey.ShouldEqual, input.expected)
			}
		})
	})
}

func TestConveyCompareNumbers(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range compareNumbersTest {
				result := false
				_, err := compareNumbers(input.guessedNumber, input.numberGenerated, input.noOfTries)
				if err == nil {
				  result = true
				}
				convey.So(result, convey.ShouldEqual, input.expectedResult)
			}
		})
	})
}

func TestConveyTakeDecision(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range compareNumbersTest {
				_, result := takeDecision(input.guessedNumber, input.numberGenerated)
				convey.So(result, convey.ShouldEqual, input.isGuessHigher)
			}
		})
	})
}

func TestConveyTriesLeft(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {

		convey.Convey("We should obtain the same value as expected value", func() {
			for _, input := range compareNumbersTest {
				_, result := printNoOfTriesLeft(input.noOfTries, input.numberGenerated)
				convey.So(result, convey.ShouldEqual, input.anyTriesLeft)
			}
		})
	})
}
