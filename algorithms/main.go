package main

import (
	"fmt"
	"time"
	"data-structures-and-algorithms-in-go/algorithms/backtracking/ratmaze"
	"data-structures-and-algorithms-in-go/algorithms/backtracking/knight-tour"
	"data-structures-and-algorithms-in-go/algorithms/backtracking/n-queen"
)

func main() {
	fmt.Println("#----------Algorithms main----------#")
	
	/* Backtracking class algorithms */
	fmt.Println("\n* 1. Backtracking *")
	//TestKnightTour(5)
	//TestRatMaze()
	TestNQueen()
}

func TestKnightTour(size int) {
	start := time.Now()
	knighttour.KnightTour(size)
	end := time.Now()
	fmt.Printf("\nElapsed time: %v\n", end.Sub(start))
}

func TestRatMaze() {
	start := time.Now()
	ratmaze.RatInAMaze()
	end := time.Now()
	fmt.Printf("\nElapsed time: %v\n", end.Sub(start))
}

func TestNQueen() {
	start := time.Now()
	queen.NQueen()
	end := time.Now()
	fmt.Printf("\nElapsed time: %v\n", end.Sub(start))
}