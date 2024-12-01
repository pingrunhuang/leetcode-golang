package trees

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	var root *TreeNode
	arr := []int{9, 2, 3, 4, 6, 78, 76, 80, 4, 3, 5}
	for _, v := range arr {
		root = root.Insert(v)
	}
	fmt.Println(kthSmallest(root, 1))

	var root2 *TreeNode
	arr2 := []int{1}
	for _, v := range arr2 {
		root2 = root2.Insert(v)
	}
	fmt.Println(kthSmallest(root2, 1))
}
