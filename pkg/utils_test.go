package pgk

import (
	"fmt"
	"testing"
)

func TestBSTInsert(*testing.T) {
	arr := []int{9, 2, 3, 4, 6, 78, 76, 80, 4, 3, 5}
	// null tree
	var tree *TreeNode
	for _, v := range arr {
		tree = tree.Insert(v)
	}
	tree.BFSView()
}

func TestBFSView(*testing.T) {
	root := TreeNode{Val: 8, Left: &TreeNode{4, &TreeNode{1, nil, nil}, nil}, Right: &TreeNode{11, &TreeNode{10, nil, nil}, &TreeNode{100, nil, nil}}}
	root.BFSView()
}

func TestReverseString(t *testing.T) {
	t1 := "abcdefg"
	e1 := "gfedcba"
	if e1 == reverseString(t1) {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
