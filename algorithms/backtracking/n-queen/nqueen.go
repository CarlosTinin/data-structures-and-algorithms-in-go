package queen

import (
	"fmt"
	
	"data-structures-and-algorithms-in-go/algorithms/backtracking/utils"
)

const QUEENS = 8

func NQueen() {
	fmt.Printf("%v Queen Solution:\n\n", QUEENS)
	board := utils.CreateBoard(QUEENS)
	positions := make([]int, QUEENS)
	
	recNqueen(positions, board, 0, 0)

	utils.PrintMatrix(QUEENS, board)
}

func isSafe(board map[int][]int, line int, column int) bool {
	for i := line; i >= 0; i-- {
		if (board[i][column] != 0) {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := line, column; i >= 0 && j >= 0 ; i, j = i-1, j-1 {
		if (board[i][j] != 0) {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := line, column; i >= 0 && j < 4 ; i, j = i-1, j+1 {
		if (board[i][j] != 0) {
			return false
		}
	}

	return true
}

func recNqueen(positions []int, board map[int][]int, line int, step int) int {
	if (step == QUEENS) {
		return 1
	}

	for column := 0; column < QUEENS; column++ {
		if (isSafe(board, line, column)) {
			board[line][column] = 1
			if (recNqueen(positions, board, line+1, step+1) == -1) {
				board[line][column] = 0
				positions[line] = 0;
			} else {
				return 1
			}
		}
	}

	return -1
}