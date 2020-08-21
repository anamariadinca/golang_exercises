package main

import "testing"

func TestNumberValid(t *testing.T) {

	inputs := []int{76, 5, 3, 10, 2, 26, 54}
  var result bool = false

	for _, input := range inputs {
		result = isNumberValid(input)
		if !result {
			t.Errorf("isNumberValid failed, invalid number: %d", input)
		}
	}

}
