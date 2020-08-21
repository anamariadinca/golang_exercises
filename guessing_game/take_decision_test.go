package main

import "testing"

func TestDecisionMaking (t *testing.T) {
  guessedNumber, numberGenerated := 14, 2

  result, expected := takeDecision(guessedNumber, numberGenerated), "Oops. Your guess was HIGH"
  if result != expected {
      t.Errorf("takeDecision failed, expected: \n %v \n received: \n %v", expected, result)
  }

  guessedNumber, numberGenerated = 1, 4
  result, expected := takeDecision(guessedNumber, numberGenerated), "Oops. Your guess was LOW"
  if result != expected  {
      t.Errorf("takeDecision failed, expected: \n %v \n received: \n %v", expected, result)
  }
}
