package pkg

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Assume a BST is defined as follows:

1. The left subtree of a node contains only nodes with keys less than or equal to the node's key.
2. The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
3. Both the left and right subtrees must also be binary search trees.
*/
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val, Left: nil, Right: nil}
		return root
	}
	switch {
	case val < root.Val:
		if root.Left == nil {
			root.Left = &TreeNode{Val: val, Left: nil, Right: nil}
			return root
		}
		root.Left.Insert(val)
	case val > root.Val:
		if root.Right == nil {
			root.Right = &TreeNode{Val: val, Left: nil, Right: nil}
			return root
		}
		root.Right.Insert(val)
	}
	return root
}

func (root *TreeNode) BFSView() {
	if root == nil {
		return
	}
	curLast := root
	nextLast := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", cur.Val)
		if cur.Left != nil {
			nextLast = cur.Left
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			nextLast = cur.Right
			queue = append(queue, cur.Right)
		}
		if curLast == cur {
			curLast = nextLast
			fmt.Println()
		}
	}
	fmt.Println()
}

func GenerateTree(values []int) *TreeNode {
	if len(values) == 1 {
		return &TreeNode{values[0], nil, nil}
	}
	var root *TreeNode
	for val := range values[1:] {
		root.Insert(val)
	}
	return root
}
