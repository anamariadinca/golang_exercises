package main

import (
	"testing"
)

func TestConversion(t *testing.T) {
	unformattedString := "5\n"
	result := convertStringToNumber(unformattedString)
	if result != 5 {
		t.Error("convertStringToNumber(\"5\n\") failed, expected: \n %v \n received: \n %v", 5, result)
	}

}
