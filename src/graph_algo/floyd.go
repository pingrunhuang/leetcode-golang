package main

import (
	"fmt"
)

/* floyd algo can only compute the shortest distance comparing path with 3 nodes at most */

type DistanceMatrix [][]float32

func initPath(distance *DistanceMatrix) *[][]int {
	n := len(*distance)
	// assume we have n poi
	path := make([][]int, n)
	for i := range *distance {
		// path[i][j] denotes the shortest path from i to j
		path[i] = make([]int, n)
		for j := range path[i] {
			path[i][j] = j
		}
	}
	return &path
}

func solve(distance *DistanceMatrix) (*[][]int, *DistanceMatrix) {
	/* O(n^3) */
	path := initPath(distance)
	numPois := len(*distance)

	for k := 0; k < numPois; k++ {
		for i := 0; i < numPois; i++ {
			for j := 0; j < numPois; j++ {
				if (*distance)[i][k]+(*distance)[k][j] < (*distance)[i][j] {
					(*distance)[i][j] = (*distance)[i][k]
					(*path)[i][j] = k
				}
			}
		}
	}
	return path, distance
}

func printDistance(distance DistanceMatrix, path [][]int) {
	n := len(distance)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				fmt.Printf("%v -> %v: %v \n", i, j, distance[i][j])
			}
		}
	}
}

func main() {
	var distance = DistanceMatrix{
		{0, 100, 100, 1.2, 9.2, 100, 0.5},
		{100, 0, 100, 5, 100, 3.1, 2},
		{100, 100, 0, 100, 100, 4, 1.5},
		{1.2, 5, 100, 0, 6.7, 100, 100},
		{9.2, 100, 100, 6.7, 0, 15.6, 100},
		{100, 3.1, 4, 100, 15.6, 0, 100},
		{0.5, 2, 1.5, 100, 100, 100, 0},
	}
	result1, result2 := solve(&distance)
	printDistance(*result2, *result1)
}
