package startgame

import (
	"fmt"
	"math/rand/v2"
)

func StartGame(level int) (string, int, error) {
	target := rand.IntN(100) + 1

	fmt.Println("Guess a number between 1 and 100")
	for i := 0; i < level; i++ {
		var guess int

		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return "", 0, err
		}

		if guess < target {
			fmt.Println("Incorrect! The number is greater than", guess, ".")
		} else if guess > target {
			fmt.Println("Incorrect! The number is less than", guess, ".")
		} else {
			fmt.Println("Congratulations! You guessed the correct number in", i+1, "attempts.", target)
			return "success", i + 1, nil
		}
	}

	fmt.Println("Sorry! You ran out of guesses. The correct number was ", target, ".")
	return "fail", level, nil
}
