package main

import (
	"fmt"

	"guessing/internal/startGame"
)

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
		startgame.StartGame(10)
	case "2", "medium":
		fmt.Println("Great! You have selected the Medium difficulty level.\nLet's start the game!")
		startgame.StartGame(5)
	case "3", "hard":
		fmt.Println("Great! You have selected the Hard difficulty level.\nLet's start the game!")
		startgame.StartGame(3)
	default:
		fmt.Println("Invalid difficulty level. Please select 1, 2, or 3.")
	}
}
