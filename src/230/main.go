package main

import (
	"fmt"

	"github.com/pingrunhuang/leetcode-golang/src/utils"
)

func getNodes(curNode *utils.TreeNode) int {
	if curNode == nil {
		return 0
	}
	return 1 + getNodes(curNode.Left) + getNodes(curNode.Right)
}

// used case
func kthSmallest(node *utils.TreeNode, k int) int {
	leftNodes := getNodes(node.Left)
	switch {
	case leftNodes < k-1:
		return kthSmallest(node.Right, k-leftNodes-1)
	case leftNodes > k-1:
		return kthSmallest(node.Left, k)
	default:
		return node.Val
	}
}

func main() {
	var root *utils.TreeNode
	arr := []int{9, 2, 3, 4, 6, 78, 76, 80, 4, 3, 5}
	for _, v := range arr {
		root = root.Insert(v)
	}
	fmt.Println(kthSmallest(root, 1))

	var root2 *utils.TreeNode
	arr2 := []int{1}
	for _, v := range arr2 {
		root2 = root2.Insert(v)
	}
	fmt.Println(kthSmallest(root2, 1))
}
