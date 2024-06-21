package ratmaze

import (
	"data-structures-and-algorithms-in-go/algorithms/backtracking/utils"
	"fmt"
)

var movei = []int{ 1, 0, -1, 0}
var movej = []int{ 0, 1, 0, -1}
const MAX_MOVES = 4

func RatInAMaze() {	
	fmt.Println("Rat in a maze")

	maze, size, i, j := readMaze()

	fmt.Println("\nLabirinto: ")	
	utils.PrintMatrix(size, maze)

	recRatInAMaze(maze, size, i, j, 0)

	fmt.Println("\nSolução: ")
	utils.PrintMatrix(size, maze)
}

func readMaze() (map[int][]int, int, int, int) {
	var size, starti, startj int
	maze := make(map[int][]int)
	
	fmt.Scan(&size)
	fmt.Scan(&starti, &startj)

	for i := 0; i < size; i++ {
		maze[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&maze[i][j])
		}
	}

	return maze, size, starti, startj
}

func recRatInAMaze(maze map[int][]int, size int, i int, j int, step int) int {
	if(maze[i][j] == 2 || step == size*size) {
		return 1
	}

	for k := 0 ; k < MAX_MOVES ; k++ {
		nexti := i + movei[k]
		nextj := j + movej[k]
		if (nexti >= 0 && nextj >= 0 && nexti < size && nextj < size) {
			if (maze[nexti][nextj] == 1) {
				maze[nexti][nextj] = 3

				if (recRatInAMaze(maze, size, nexti, nextj, step+1) == -1) {
					maze[nexti][nextj] = 1
				} else {
					return 1
				}
			} else if (maze[nexti][nextj] == 2) {
				maze[nexti][nextj] = 9
				return 1
			}
		}
	}

	return -1
}