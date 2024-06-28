package queen

import (
	"fmt"
	
	"data-structures-and-algorithms-in-go/algorithms/backtracking/utils"
)

const QUEENS = 16

func NQueen() {
	fmt.Printf("%v Queen Solution:\n\n", QUEENS)
	board := utils.CreateBoard(QUEENS)

	diagonals := (4*QUEENS - 2)/2
	leftDiag  := make([]int, diagonals)
	rightDiag := make([]int, diagonals)
	checkColumn := make([]int, QUEENS)
	
	//recNqueen(board, 0, 0)
	recNqueenEnhanced(board, 0, 0, leftDiag, rightDiag, checkColumn)

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

func recNqueenEnhanced(board map[int][]int, line int, step int, leftDiag []int, rightDiag []int, checkColumn []int) int {
	if (step == QUEENS) {
		return 1
	}

	for column := 0; column < QUEENS; column++ {
		if (rightDiag[line+column] == 0 && leftDiag[line-column+QUEENS-1] == 0 && checkColumn[column] == 0) {
			board[line][column] = 1
			rightDiag[line+column], leftDiag[line-column+QUEENS-1], checkColumn[column] = 1, 1, 1
			if (recNqueen(board, line+1, step+1) == -1) {
				board[line][column] = 0
				rightDiag[line+column], leftDiag[line-column+QUEENS-1], checkColumn[column] = 0, 0, 0
			} else {
				return 1
			}
		}
	}

	return -1
}

func recNqueen(board map[int][]int, line int, step int) int {
	if (step == QUEENS) {
		return 1
	}

	for column := 0; column < QUEENS; column++ {
		if (isSafe(board, line, column)) {
			board[line][column] = 1
			if (recNqueen(board, line+1, step+1) == -1) {
				board[line][column] = 0
			} else {
				return 1
			}
		}
	}

	return -1
}