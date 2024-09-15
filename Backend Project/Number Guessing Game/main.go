package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	highScores :=map[string]int {
		"Easy": 10,
		"Medium": 5,
		"Hard": 3,
	}

	for {
		playGame(highScores)

		fmt.Print("\nDo you want to play again? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.TrimSpace(playAgain)

		if strings.ToLower(playAgain) != "y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func playGame(highScores map[string]int) {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min+1) + min
	// fmt.Println("The secret number is:", secretNumber) // For debugging purposes

	fmt.Println("\nWelcome to the Number Guessing Game!")
	fmt.Println("Guess a number between 1 and 100")

	reader := bufio.NewReader(os.Stdin)
	var maxAttempts int
	var difficulty string

	for {
		fmt.Println("\nDifficulty Levels")
		fmt.Println("1 - Easy (10 attempts)")
		fmt.Println("2 - Medium (5 attempts)")
		fmt.Println("3 - Hard (3 attempts)")
		fmt.Print("Choose a difficulty level between 1 to 3: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			maxAttempts = 10
			difficulty = "Easy"
			fmt.Println("You have chosen Easy difficulty. You have 10 attempts.")
		case "2":
			maxAttempts = 5
			difficulty = "Medium"
			fmt.Println("You have chosen Medium difficulty. You have 5 attempts.")
		case "3":
			maxAttempts = 3
			difficulty = "Hard"
			fmt.Println("You have chosen Hard difficulty. You have 3 attempts.")
		default:
			fmt.Println("Invalid input. Please select 1 for Easy, 2 for Medium, or 3 for Hard.")
			continue
		}
		break
	}

	startTime := time.Now()
	attempts := 0
	hintsGiven := 0

	for attempts < maxAttempts {
		attempts++
		fmt.Print("\nPlease input your guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			return
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again.")
		} else {
			elapsed := time.Since(startTime)
			fmt.Println("Congratulations! You guessed the secret number in", attempts, "attempts and", elapsed.Seconds(), "seconds.")

			if attempts < highScores[difficulty] {
				fmt.Printf("New high score for %s difficulty!\n", difficulty)
				highScores[difficulty] = attempts
			}

			fmt.Printf("Current high score for %s difficulty: %d attempts\n", difficulty, highScores[difficulty])
			return
		}

		// Hint system after 2 failed attempts
		if attempts > 1 && hintsGiven == 0 {
			if guess < secretNumber {
				fmt.Println("Hint: Try a larger number.")
			} else {
				fmt.Println("Hint: Try a smaller number.")
			}
			hintsGiven++
		}

		if attempts == maxAttempts {
			fmt.Println("You have reached the maximum number of attempts. The secret number was:", secretNumber)
		}
	}
}
