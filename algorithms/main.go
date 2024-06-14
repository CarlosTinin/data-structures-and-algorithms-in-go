package main

import (
	"fmt"
	"time"

	"data-structures-and-algorithms-in-go/algorithms/backtracking/knight-tour"
)

func main() {
	fmt.Println("#----------Algorithms main----------#")
	
	/* Backtracking class algorithms */
	fmt.Println("\n* 1. Backtracking *")
	start := time.Now()
	knighttour.KnightTour(5)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Elapsed time: %v", elapsed)
	fmt.Println()
}