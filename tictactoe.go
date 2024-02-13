package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board [3][3]string
var currentPlayer string

const (
	emptyCell = " "
	playerX   = "X"
	playerO   = "O"
)

func main() {
	initBoard()
	printBoard()

	for {
		if currentPlayer == playerX {
			makeMove()
		} else {
			minimaxMove()
		}

		printBoard()

		if checkWin() {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		if checkDraw() {
			fmt.Println("It's a draw!")
			break
		}

		switchPlayer()
	}
}

func initBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = emptyCell
		}
	}
	currentPlayer = playerX
}

func printBoard() {
	clearScreen()
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----")
		}
	}
	fmt.Println()
}

func makeMove() {
	var row, col int

	for {
		fmt.Printf("Player %s, enter your move (row and column): ", currentPlayer)
		fmt.Scan(&row, &col)

		if isValidMove(row, col) {
			board[row][col] = currentPlayer
			break
		} else {
			fmt.Println("Invalid move, try again.")
		}
	}
}

func minimaxMove() {
	bestScore := -1000
	bestMove := Move{-1, -1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == emptyCell {
				board[i][j] = currentPlayer
				score := minimax(board, 0, false)
				board[i][j] = emptyCell

				if score > bestScore {
					bestScore = score
					bestMove = Move{i, j}
				}
			}
		}
	}

	if bestMove.Row != -1 && bestMove.Col != -1 {
		board[bestMove.Row][bestMove.Col] = currentPlayer
	}
}

func minimax(board [3][3]string, depth int, isMaximizing bool) int {
	if checkWin() {
		if currentPlayer == playerX {
			return -1
		} else {
			return 1
		}
	}

	if checkDraw() {
		return 0
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == emptyCell {
					board[i][j] = playerX
					score := minimax(board, depth+1, false)
					board[i][j] = emptyCell
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == emptyCell {
					board[i][j] = playerO
					score := minimax(board, depth+1, true)
					board[i][j] = emptyCell
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

func isValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != emptyCell {
		return false
	}
	return true
}

func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}

	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}

	return false
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == emptyCell {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == playerX {
		currentPlayer = playerO
	} else {
		currentPlayer = playerX
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Move struct {
	Row, Col int
}
