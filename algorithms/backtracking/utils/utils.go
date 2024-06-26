package utils

import (
	"fmt"
)

func PrintMatrix(n int, board map[int][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v ", board[i][j])
		}
		fmt.Println()
	}
}

func CreateBoard(size int) map[int][]int {
	board := make(map[int][]int)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return board;
}