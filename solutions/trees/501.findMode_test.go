package trees

import (
	"fmt"
	"testing"

	"github.com/leetcode-golang/pkg"
)

func TestFindMode(t *testing.T) {
	var test_cases []int = []int{1, 2, 2}
	root := pkg.GenerateTree(test_cases)
	result := findMode(root)
	fmt.Printf("t: %v\n", t)
	fmt.Print((result))
}
