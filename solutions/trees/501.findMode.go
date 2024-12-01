package trees

import (
	"math"
)

/**
 * 501. Find Mode in Binary Search Tree
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// a slice of pair that implements the sort interface so that you can use the sort function
// type pairList []pair

// func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// func (p pairList) Less(i, j int) bool { return p[i].val < p[j].val }

// func (p pairList) Len() int { return len(p) }

func dfs(node *TreeNode, mapper map[int]int) {

	if node != nil {
		mapper[node.Val]++
		if node.Left != nil {
			dfs(node.Left, mapper)
		}
		if node.Right != nil {
			dfs(node.Right, mapper)
		}
	}
}

func findMode(root *TreeNode) []int {
	mapper := make(map[int]int)
	result := make([]int, 0)
	dfs(root, mapper)
	MAXINT := math.MinInt64
	for k, v := range mapper {
		if v >= MAXINT {
			if v > MAXINT {
				MAXINT = v
				result = result[0:0]
			}
			result = append(result, k)
		}
	}
	return result
}
