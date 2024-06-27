package knighttour

import (
	"fmt"
	
	"data-structures-and-algorithms-in-go/algorithms/backtracking/utils"
)

var movei = []int{ 2, 1, -1, -2, -2, -1, 1, 2}
var movej = []int{ 1, 2, 2, 1, -1, -2, -2, -1}
const MAX_MOVES = 8

/*
	Knight Tour Auxiliary main function, starts the recursive backtraking function and prints the matrix result

	Parameter: n int size of the board 
*/
func KnightTour(n int) {
	fmt.Printf("%v Knight Tour Solution\n\n", n)
	board := utils.CreateBoard(n)
	board[0][0] = -1

	recKnightTour(n, board, 0, 0, 1)
	utils.PrintMatrix(n, board)
}


func recKnightTour(n int, board map[int][]int, i int, j int, step int) int {
	if (step == n*n) {
		return 1
	}

	for k := 0; k < MAX_MOVES; k++ {
		nexti := i + movei[k]
		nextj := j + movej[k]
		if ( nexti >= 0 && nextj >= 0 && nexti < n && nextj < n) {			
			if (board[ nexti ][ nextj ] == 0) {
				board[ nexti ][ nextj ] = step
				
				if ( recKnightTour(n, board, nexti, nextj, step+1) == -1 ) {
					board[ nexti ][ nextj ] = 0
				} else {
					return 1
				}
			}
		}
	}

	return -1
}