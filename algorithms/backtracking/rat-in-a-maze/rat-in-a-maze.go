package ratmaze

import (
	"data-structures-and-algorithms-in-go/algorithms/backtracking/utils"
	"fmt"
)

func RatInAMaze() {	
	fmt.Println("Rat in a maze")
	maze, size, _, _ := readMaze()
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

// func readMaze(string fileName) map[int][]int {
// 	file, err := os.Open("entries/" + fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	reader := bufio.NewReader(file)

// 	for {
// 		char, err := reader.ReadByte()
// 		if err != nil {
// 			if err.Error() == "EOF" {
// 				break
// 			}
// 			log.Fatal(err)
// 		}
		
// 	}
// }

// func recRatInAMaze() {
// 	if(true) {

// 	}

// 	for {

// 	}
// }