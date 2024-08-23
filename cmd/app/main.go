package main

import (
	"fmt"
	"math/rand/v2"

	"C:\Users\Т\Desktop\Guessing-game\internal"
)



func startGame(level int) {
	target := rand.IntN(100) + 1

	var records []GameRecord

	fmt.Println("Guess a number between 1 and 100")
	for i := 0; i < level; i++ {

		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}

		if guess < target {
			fmt.Println("Incorrect! The number is greater than", guess, ".")
		} else if guess > target {
			fmt.Println("Incorrect! The number is less than", guess, ".")
		} else {
			fmt.Println("Congratulations! You guessed the correct number in", i+1, "attempts.")
			records = append(records, GameRecord{Result: "success", Attempts: i + 1})
			err := saveToCSV(records, "records.csv")
			if err != nil {
				fmt.Println("Failed to save records to CSV:", err)
			}
			return
		}
	}

	fmt.Println("Sorry! You ran out of guesses. The correct number was ", target, ".")
	records = append(records, GameRecord{Result: "fail", Attempts: level})
	err := saveToCSV(records, "records.csv")
	if err != nil {
		fmt.Println("Failed to save records to CSV:", err)
	}
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number")

	fmt.Println("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

	var levelInput string
	_, err := fmt.Scanln(&levelInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter 1, 2, or 3.")
	}
	switch levelInput {
	case "1", "easy":
		fmt.Println("Great! You have selected the Easy difficulty level.\nLet's start the game!")
		startGame(10)
	case "2", "medium":
		fmt.Println("Great! You have selected the Medium difficulty level.\nLet's start the game!")
		startGame(5)
	case "3", "hard":
		fmt.Println("Great! You have selected the Hard difficulty level.\nLet's start the game!")
		startGame(3)
	default:
		fmt.Println("Invalid difficulty level. Please select 1, 2, or 3.")
	}
}
