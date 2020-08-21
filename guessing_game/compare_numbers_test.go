package main

import (
  "testing"
)

type inputTest struct {
  guessedNumber int
  numberGenerated int
  noOfTries int
}

func NewInput(guessedNumber int, numberGenerated int, noOfTries int) inputTest {
  in := inputTest{guessedNumber, numberGenerated, noOfTries}
  return in
}

func TestCompareNumbers (t *testing.T) {

  i1 := NewInput(14,2,3)
  i2 := NewInput(4,5,5)
  i3 := NewInput(6,1,9)
  i4 := NewInput(40,54,2)
  i5 := NewInput(23,98,7)

  inputList := []inputTest{i1, i2, i3, i4, i5}

  for _, in := range inputList {
    result := compareNumbers(in.guessedNumber, in.numberGenerated, in.noOfTries)
    expectedResult := takeDecision(in.guessedNumber, in.numberGenerated) + "\n" + printNoOfTriesLeft(in.noOfTries+1, in.numberGenerated)
    if result != expectedResult{
      t.Errorf("compareNumbers failed, expected: \n %v \n received: \n %v \n guessedNumber: %d \n numberGenerated: %d", expectedResult, result, in.guessedNumber, in.numberGenerated)
    }
  }

}
