package arrays

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	// how to make 2d matrix in golang
	test1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	result1 := minPathSum(test1)
	if result1 == 7 {
		fmt.Printf("Test case 1 passed\n")
	} else {
		fmt.Printf("Test case 1 failed\n")
	}
}
