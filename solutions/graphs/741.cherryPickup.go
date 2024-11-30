package graphs

import "fmt"

/*In a N x N grid representing a field of cherries, each cell is one of three possible integers.

0 means the cell is empty, so you can pass through;
1 means the cell contains a cherry, that you can pick up and pass through;
-1 means the cell contains a thorn that blocks your way.
Your task is to collect maximum number of cherries possible by following the rules below:

Starting at the position (0, 0) and reaching (N-1, N-1) by moving right or down through valid path cells (cells with value 0 or 1);
After reaching (N-1, N-1), returning to (0, 0) by moving left or up through valid path cells;
When passing through a path cell containing a cherry, you pick it up and the cell becomes an empty cell (0);
If there is no valid path between (0, 0) and (N-1, N-1), then no cherries can be collected.
Example 1:
Input: grid =
[[0, 1, -1],
 [1, 0, -1],
 [1, 1,  1]]
Output: 5
Explanation:
The player started at (0, 0) and went down, down, right right to reach (2, 2).
4 cherries were picked up during this single trip, and the matrix becomes [[0,1,-1],[0,0,-1],[0,0,0]].
Then, the player went left, up, up, left to return home, picking up one more cherry.
The total number of cherries picked up is 5, and this is the maximum possible.
Note:

grid is an N by N 2D array, with 1 <= N <= 50.
Each grid[i][j] is an integer in the set {-1, 0, 1}.
It is guaranteed that grid[0][0] and grid[N-1][N-1] are not -1.
*/

/*
Solution 1: imagine 2 people walking from top to down at the same time.
*/

func cherryPickup1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	// how to create a grid with the same size as the parameters
	var sum [][]int
	// copy()
	for i := 0; i < rows; i++ {
		sum[i] = grid[i]
	}

	copy(grid, sum)
	fmt.Println(&sum == &grid)

	// first trip from top left to down right
	for row := rows - 1; row >= 1; row-- {
		// boundary setting
		if grid[row-1][0] == -1 {
			sum[row][0] = 0
		} else {
			sum[row][0] += grid[row-1][0]
			grid[row][0] = 0
		}
		for col := cols - 1; col >= 1; col-- {
			if grid[row][col-1] == -1 && grid[row-1][col] == -1 {
				sum[row][col] = 0
			} else if grid[row][col-1] == -1 {
				sum[row][col] += grid[row-1][col]
				grid[row][col] = 0
			} else if grid[row-1][col] == -1 {
				sum[row][col] += grid[row][col-1]
				grid[row][col] = 0
			} else if grid[row][col-1] < grid[row-1][col] {
				sum[row][col] += grid[row-1][col]
				grid[row][col] = 0
			} else {
				sum[row][col] += grid[row][col-1]
				grid[row][col] = 0
			}
		}
	}

	// row := rows - 1
	// col := cols - 1
	// for row >= 0 && col >= 0 {
	// 	if row == 0 &&
	// }

	// back trip from down right to top left
	// for row := 0; row < rows; row++ {
	// 	// boundary setting
	// 	if grid[row-1][0] == -1 {
	// 		sum[row][0] = 0
	// 	} else {
	// 		sum[row][0] = grid[row-1][0]
	// 		grid[row-1][0] = 0
	// 	}
	// 	for col := cols; col >= 1; col-- {
	// 		if grid[row][col-1] == -1 && grid[row-1][col] == -1 {
	// 			sum[row][col] = 0
	// 		} else if grid[row][col-1] == -1 {
	// 			sum[row][col] = grid[row-1][col]
	// 			grid[row-1][col] = 0
	// 		} else if grid[row-1][col] == -1 {
	// 			sum[row][col] = grid[row][col-1]
	// 			grid[row][col-1] = 0
	// 		} else if grid[row][col-1] < grid[row-1][col] {
	// 			sum[row][col] = grid[row-1][col]
	// 			grid[row-1][col] = 0
	// 		} else {
	// 			sum[row][col] = grid[row][col-1]
	// 			grid[row][col-1] = 0
	// 		}
	// 	}
	// }
	fmt.Println(grid)
	fmt.Println(sum)
	return 0
}

func cherryPickup2(grid [][]int) {
	rows := len(grid)
	// cols := len(grid[0])
	// how to create a grid with the same size as the parameters
	var sum = make([][]int, rows)
	for x := 0; x < rows; x++ {
		// for y := 0; y < cols; y++ {
		// 	sum[x][y] = grid[x][y]
		// }
		sum[x] = grid[x]
	}
	// sum[0] = []int{1, 2, 3}
	sum[0][0] = 4
	fmt.Println(sum)
	fmt.Println(grid)
	// fmt.Printf("%p\n", &sum[0])
	// fmt.Printf("%p\n", &grid[0])
}

func copyMatrix(srcMatrix [][]int) [][]int {
	copyGrid := make([][]int, len(srcMatrix))
	for i := 0; i < len(srcMatrix); i++ {
		copyGrid[i] = make([]int, len(srcMatrix[0]))
		for j := 0; j < len(srcMatrix[0]); j++ {
			copyGrid[i][j] = srcMatrix[i][j]
		}
	}
	return copyGrid
}
