package main

import(
  "testing"
  "strconv"
)

func TestTriesLeftOutput (t *testing.T) {
  noOfTries, numberGenerated := 3, 15

  result := printNoOfTriesLeft(noOfTries, numberGenerated)
  expectedResult := "You have " + strconv.Itoa(10 - noOfTries - 1) + " tries left"
  if result != expectedResult {
    t.Errorf("printNoOfTriesLeft failed, expected: \n %v \n received: \n %v ", expectedResult, result)
  }

  noOfTries, numberGenerated = 9, 25
  result = printNoOfTriesLeft(noOfTries, numberGenerated)
  expectedResult = "Sorry. You didn't guess my number. It was: " + strconv.Itoa(numberGenerated)
  if result != expectedResult {
    t.Errorf("printNoOfTriesLeft failed, expected: \n %v \n received: \n %v ", expectedResult, result)
  }
}
