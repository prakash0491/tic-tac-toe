package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var board [3][3]string
var currentPlayer string

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func displayBoard() {
	clearScreen()
	fmt.Println("  1 2 3")
	fmt.Println(" -------")
	for i := 0; i < 3; i++ {
		fmt.Printf("|%s|%s|%s|\n", board[i][0], board[i][1], board[i][2])
		fmt.Println(" -------")
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
		if board[0][i] != " " && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}

func main() {
	initializeBoard()
	displayBoard()
	currentPlayer = "X"

	for {
		var row, col int

		fmt.Printf("Player %s's turn. Enter row (1-3) and column (1-3) separated by space: ", currentPlayer)
		_, err := fmt.Scanln(&row, &col)
		if err != nil || row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != " " {
			fmt.Println("Invalid move! Try again.")
			continue
		}

		board[row-1][col-1] = currentPlayer
		displayBoard()

		if checkWin() {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		if strings.Contains(fmt.Sprint(board), " ") == false {
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
