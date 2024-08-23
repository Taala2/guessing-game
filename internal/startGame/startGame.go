package startgame

import (
	"fmt"
	"math/rand/v2"

	"guessing/internal/history"
)

func StartGame(level int) {
	target := rand.IntN(100) + 1

	var records []history.GameRecord

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
			records = append(records, history.GameRecord{Result: "success", Attempts: i + 1})
			err := history.SaveToCSV(records, "records.csv")
			if err != nil {
				fmt.Println("Failed to save records to CSV:", err)
			}
			return
		}
	}

	fmt.Println("Sorry! You ran out of guesses. The correct number was ", target, ".")
	records = append(records, history.GameRecord{Result: "fail", Attempts: level})
	err := history.SaveToCSV(records, "records.csv")
	if err != nil {
		fmt.Println("Failed to save records to CSV:", err)
	}
}
