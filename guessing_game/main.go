package main

import (
	"math/rand"
	"strconv"
	"fmt"
	"time"
)

func main() {
	fmt.Println("I'll pick a number from 1 to 100. Try to guess it. You have 10 tries.")
	noOfTries, numberWasGuessed := 0, false
	rand.Seed(time.Now().UnixNano())
	numberGenerated := rand.Intn(100) + 1

	for !numberWasGuessed {
		var guessedNumber int
		_, err:= fmt.Scan(&guessedNumber)
		if err != nil {
			fmt.Errorf("blah")
			return
		}
		if guessedNumber == numberGenerated {
			fmt.Println("Good job, you guessed it!")
			numberWasGuessed = true
		} else {
			fmt.Println(compareNumbers(guessedNumber, numberGenerated, noOfTries))
		}
	}

}

func compareNumbers(guessedNumber int, numberGenerated int, noOfTries int) (string, error) {
	if isNumberValid(guessedNumber) {
		noOfTries++
		takeDecisionMessage, _ := takeDecision(guessedNumber, numberGenerated)
		printNoOfTriesLeftMessage, _ := printNoOfTriesLeft(noOfTries, numberGenerated)
		return takeDecisionMessage + "\n" + printNoOfTriesLeftMessage, nil
	} else {
		return "", fmt.Errorf("Please introduce a valid number; between 1 and 100")
	}
}

func printNoOfTriesLeft(noOfTries int, numberGenerated int) (string, bool) {
	if 10-noOfTries-1 > 0 {
		return "You have " + strconv.Itoa(10 - noOfTries - 1) + " tries left", true
	} else {
		return "Sorry. You didn't guess my number. It was: " + strconv.Itoa(numberGenerated), false
	}
}

func isNumberValid(guessedNumber int) bool {
	if guessedNumber <= 100 && guessedNumber >= 1 {
		return true
	}
	return false
}

func takeDecision(guessedNumber int, numberGenerated int) (string, bool) {
	if guessedNumber < numberGenerated {
		return "Oops. Your guess was LOW", false

	} else {
		return "Oops. Your guess was HIGH", true
	}

}
