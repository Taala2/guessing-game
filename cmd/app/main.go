package main

import (
	"fmt"

	"guessing/internal/history"
	"guessing/internal/startGame"
)


var levelInput string
var records []history.GameRecord

var level int
var filepath = "../../internal/data/records.csv"

func main() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.")

	fmt.Println("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

	commands()

}

func commands() {
	_, err := fmt.Scanln(&levelInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter 1, 2, or 3.")
	}
	switch levelInput {
	case "1", "easy":
		fmt.Println("Great! You have selected the Easy difficulty level.\nLet's start the game!")
		level = 10
	case "2", "medium":
		fmt.Println("Great! You have selected the Medium difficulty level.\nLet's start the game!")
		level = 5
	case "3", "hard":
		fmt.Println("Great! You have selected the Hard difficulty level.\nLet's start the game!")
		level = 3

	case "list":
		stats, _ := history.ReadCSV(filepath)
		fmt.Println("Game History:\nResult, 	Attempts")
		for _, stat := range stats {
			fmt.Printf("%s, 	%d\n", stat.Result, stat.Attempts)
		}
		commands()
	default:
		fmt.Println("Invalid difficulty level. Please select 1, 2, or 3.")
		commands()
	}

	end, i, _ := startgame.StartGame(level)

	records = append(records, history.GameRecord{Result: end, Attempts: i})
	history.SaveToCSV(records, filepath)
}