# Number Guessing Game (CLI)

## Overview
This is a Command-Line Interface (CLI) based game where the user tries to guess a randomly selected number between 1 and 100. The game provides various difficulty levels that determine how many chances the user has to guess the correct number. After each guess, the game will provide feedback on whether the guessed number is higher or lower than the target number. The game continues until the user guesses the correct number or runs out of chances. 

To make the game more engaging, several additional features have been implemented, including multiple rounds, a timer, a hint system, and high score tracking.

## How to Play
1. When you start the game, you will be greeted with a welcome message and the rules.
2. The computer will randomly select a number between 1 and 100.
3. You will be prompted to select a difficulty level:
   - **Easy**: More chances to guess the number.
   - **Medium**: Fewer chances than Easy.
   - **Hard**: The fewest chances to guess the number.
4. Enter your guess.
5. After each guess:
   - If the guess is correct, the game will display a congratulatory message and show how many attempts it took you to guess correctly.
   - If the guess is incorrect, the game will tell you whether the target number is higher or lower than your guess.
6. The game ends when you guess the correct number or run out of chances.

### Additional Features
- **Multiple Rounds**: After each game, you can choose to play again.
- **Timer**: A timer records how long it takes to guess the number.
- **Hint System**: Hints may be provided if you're stuck.
- **High Score Tracking**: The game keeps track of the fewest attempts it took to guess the number for each difficulty level.

## Instructions to Run
1. Clone the repository or download the game files.
2. Open your terminal or command prompt.
3. Navigate to the directory where the game is located.
4. Run the game using the following command:
   ```bash
   go run main.go 
5. Follow the on-screen instructions to start playing.

Happy Coding!