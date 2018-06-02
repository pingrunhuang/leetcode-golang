package leetcode

import (
	"fmt"
	"testing"
)

func TestCherryPickup(t *testing.T) {
	case1 := [][]int{
		[]int{0, 1, -1},
		[]int{1, 0, -1},
		[]int{1, 1, 1},
	}
	// cherryPickup2(case1)
	copyCase1 := copyMatrix(case1)

	copyCase1[0][0] = 12
	fmt.Println(case1)
	fmt.Println(copyCase1)
}
