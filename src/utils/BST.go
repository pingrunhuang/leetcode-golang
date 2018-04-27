package utils

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func (root *Tree) Insert(val int) *Tree {
	if root == nil {
		root = &Tree{Val: val, Left: nil, Right: nil}
		return root
	}
	switch {
	case val < root.Val:
		if root.Left == nil {
			root.Left = &Tree{Val: val, Left: nil, Right: nil}
			return root
		}
		root.Left.Insert(val)
	case val > root.Val:
		if root.Right == nil {
			root.Right = &Tree{Val: val, Left: nil, Right: nil}
			return root
		}
		root.Right.Insert(val)
	}
	return root
}

func (root *Tree) BFSView() {
	if root == nil {
		return
	}
	curLast := root
	nextLast := root
	queue := make([]*Tree, 0)
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
