package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodes(curNode *TreeNode) int {
	if curNode == nil {
		return 0
	}
	return 1 + getNodes(curNode.Left) + getNodes(curNode.Right)
}

// used case
func kthSmallest(node *TreeNode, k int) int {
	leftNodes := getNodes(node.Left)
	switch {
	case leftNodes < k:
		return kthSmallest(node.Right, k-leftNodes-1)
	case leftNodes > k:
		return kthSmallest(node.Left, k)
	default:
		return node.Val
	}
}

func main() {
}
