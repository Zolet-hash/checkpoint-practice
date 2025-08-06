package main

import "fmt"

// EightQueens prints all solutions to the 8 queens puzzle using recursion.
func EightQueens() {
	var board [8]int
	solve(0, &board)
}

// solve places queens recursively starting from the given column.
func solve(col int, board *[8]int) {
	if col == 8 {
		printSolution(*board)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

// isSafe checks if placing a queen at (col, row) is safe.
func isSafe(col int, row int, board *[8]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		// Check same row
		if prevRow == row {
			return false
		}

		// Check diagonals
		if abs(prevRow-row) == abs(prevCol-col) {
			return false
		}
	}
	return true
}

// printSolution prints the board as a single 8-digit number.
func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		fmt.Print(board[i])
	}
	fmt.Println()
}

// abs returns the absolute value of an int.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	EightQueens()
}
