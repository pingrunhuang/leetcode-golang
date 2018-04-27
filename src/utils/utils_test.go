package utils

import (
	"testing"
)

func TestBSTInsert(*testing.T) {
	arr := []int{9, 2, 3, 4, 6, 78, 76, 80, 4, 3, 5}
	// null tree
	var tree *Tree
	for _, v := range arr {
		tree = tree.Insert(v)
	}
	tree.BFSView()
}

func TestBFSView(*testing.T) {
	root := Tree{Val: 8, Left: &Tree{4, &Tree{1, nil, nil}, nil}, Right: &Tree{11, &Tree{10, nil, nil}, &Tree{100, nil, nil}}}
	root.BFSView()
}
