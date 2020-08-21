package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := "26\n"
	reader := bufio.NewReader(strings.NewReader(input))

	value := readFromKeyboard(reader)

	if value != input {
		t.Errorf("readFromKeyboard failed, expected: %v,\n received: %v", input, value)
	}
}
