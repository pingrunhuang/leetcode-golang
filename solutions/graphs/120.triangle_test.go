package graphs

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	t1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	minimumTotal(t1)
	// if r1 == 11 {
	// 	fmt.Println("Passed")
	// } else {
	// 	fmt.Printf("Expected 11 but %d\n", r1)
	// }

	// t2 := [][]int{{-1}, {2, 3}, {1, -1, -3}}
	// r2 := minimumTotal(t2)

	// if r2 == -1 {
	// 	fmt.Println("Passed")
	// } else {
	// 	fmt.Printf("Expected -1 but %d\n", r2)
	// }
}
