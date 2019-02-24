package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
