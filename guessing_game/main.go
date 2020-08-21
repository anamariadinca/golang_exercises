package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("I'll pick a number from 1 to 100. Try to guess it. You have 10 tries.")
	noOfTries, numberWasGuessed := 0, false
	rand.Seed(time.Now().UnixNano())
	numberGenerated := rand.Intn(100) + 1
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	for !numberWasGuessed {
		input := readFromKeyboard(reader)
		guessedNumber := convertStringToNumber(input)
		if guessedNumber == numberGenerated {
			fmt.Println("Good job, you guessed it!")
			numberWasGuessed = true
		} else {
			fmt.Println(compareNumbers(guessedNumber, numberGenerated, noOfTries))
		}
	}

}

func compareNumbers(guessedNumber int, numberGenerated int, noOfTries int) string {
	if isNumberValid(guessedNumber) {
		noOfTries++
		return takeDecision(guessedNumber, numberGenerated) + "\n" + printNoOfTriesLeft(noOfTries, numberGenerated)
	} else {
		return "Please introduce a valid number; between 1 and 100"
	}
}

func printNoOfTriesLeft(noOfTries int, numberGenerated int) string {
	if 10-noOfTries-1 > 0 {
		return "You have " + strconv.Itoa(10 - noOfTries - 1) + " tries left"
	} else {
		return "Sorry. You didn't guess my number. It was: " + strconv.Itoa(numberGenerated)
	}
}

func isNumberValid(guessedNumber int) bool {
	if guessedNumber <= 100 && guessedNumber >= 1 {
		return true
	}
	return false
}

func takeDecision(guessedNumber int, numberGenerated int) string {
	if guessedNumber < numberGenerated {
		return "Oops. Your guess was LOW"

	} else {
		return "Oops. Your guess was HIGH"
	}

}

func convertStringToNumber(s string) int {
	s = strings.TrimSpace(s)
	guessedNumber, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return guessedNumber
}

func readFromKeyboard(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}
