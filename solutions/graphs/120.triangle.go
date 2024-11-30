package graphs

import (
	"strconv"
)

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (38.22%)
 * Total Accepted:    168.3K
 * Total Submissions: 440.2K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 * Category: DP
 */

// notice for col j on row i, the next step can only be col j or col j+1 in row i+1
func minimumTotal(triangle [][]int) int {
	memo := make(map[string]int)
	totalRows := len(triangle)
	// last_cols := len(triangle)
	for j, v := range triangle[totalRows-1] {
		key := strconv.Itoa(totalRows-1) + strconv.Itoa(j)
		memo[key] = v
	}
	for i := totalRows - 2; i >= 0; i-- {
		for j, v := range triangle[i] {
			key := strconv.Itoa(i) + strconv.Itoa(j)
			if memo[strconv.Itoa(i+1)+strconv.Itoa(j)] > memo[strconv.Itoa(i+1)+strconv.Itoa(j+1)] {
				memo[key] = v + memo[strconv.Itoa(i+1)+strconv.Itoa(j+1)]
			} else {
				memo[key] = v + memo[strconv.Itoa(i+1)+strconv.Itoa(j)]
			}
		}
	}
	return memo[strconv.Itoa(0)+strconv.Itoa(0)]
}
