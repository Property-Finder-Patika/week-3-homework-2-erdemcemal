package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

var numberOfRightPositions int
var numberOfWrongPositions int

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Guess a number between 1 and 9999")
	fmt.Println("You have only 5 tries to guess the number")
	fmt.Println("Good luck!")

	numberOfDigits := 4
	number, err := generateRandomNumber(numberOfDigits)
	if err != nil {
		fmt.Println("Error generating random number")
		return
	}
	fmt.Println("The number is: ", number)
	for i := 0; i < 5; i++ {
		fmt.Printf("Guess #%d: ", i+1)
		var guess int

		fmt.Scanf("%d", &guess)
		numberOfRightPositions, numberOfWrongPositions = compareNumbers(number, guess)
		if numberOfRightPositions == 4 {
			fmt.Println("You guessed the number!")
			break
		} else {
			fmt.Printf("You have %d right positions and %d wrong positions\n", numberOfRightPositions, numberOfWrongPositions)
		}
		// reset the number of right and wrong positions
		numberOfRightPositions = 0
		numberOfWrongPositions = 0
	}
}

func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int(math.Pow(10, float64(numberOfDigits)) - 1)
	minLimit := int(math.Pow(10, float64(numberOfDigits-1)))

	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(maxLimit)))
	if err != nil {
		return 0, err
	}
	number := int(randNumber.Int64())

	if number <= minLimit {
		number += minLimit
	}

	if number > maxLimit {
		number = maxLimit
	}
	return number, nil
}

func compareNumbers(n, g int) (int, int) {
	number := strconv.Itoa(n)
	guess := strconv.Itoa(g)

	if len(number) != len(guess) {
		fmt.Println("Your guess must have 4 digits")
		os.Exit(1)
	}

	dict := make(map[string]int)
	for i := 0; i < 4; i++ {
		dict[string(number[i])] = i
	}
	for i := 0; i < 4; i++ {
		value := guess[i]
		if position, ok := dict[string(value)]; ok {
			if position == i {
				numberOfRightPositions++
			} else {
				numberOfWrongPositions++
			}
		}

	}
	return numberOfRightPositions, numberOfWrongPositions
}
