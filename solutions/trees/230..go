package trees

/*
230.
*/

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
	case leftNodes < k-1:
		return kthSmallest(node.Right, k-leftNodes-1)
	case leftNodes > k-1:
		return kthSmallest(node.Left, k)
	default:
		return node.Val
	}
}
