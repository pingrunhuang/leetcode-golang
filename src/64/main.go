package main

import "fmt"

// DP
// Minimum Path Sum
// Given a m x n grid filled with non-negative numbers,
// find a path from top left to bottom right which minimizes the sum of all numbers along its path.
// Example:
// [[1,3,1],
//  [1,5,1],
//  [4,2,1]]
// return 7. Because 1->3->1->1->1 minimize the sum

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	// sum for maintaining the minimum path
	sum := grid

	// boundary setting
	// first column is always the
	for i := 1; i < rows; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}

	for j := 1; j < cols; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if sum[i][j-1] < sum[i-1][j] {
				sum[i][j] = sum[i][j-1] + grid[i][j]
			} else {
				sum[i][j] = sum[i-1][j] + grid[i][j]
			}
		}
	}
	return sum[rows-1][cols-1]
}

//
func Launch64() {
	// how to make 2d matrix in golang
	test1 := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	result1 := minPathSum(test1)
	if result1 == 7 {
		fmt.Printf("Test case 1 passed\n")
	} else {
		fmt.Printf("Test case 1 failed\n")
	}
}

func main() {

}
