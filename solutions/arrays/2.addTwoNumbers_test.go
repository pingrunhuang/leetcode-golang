package arrays

import "testing"

func TestAddTwoNum(t *testing.T) {
	// x1 := ListNode{Val: 8, Next: nil}
	// x2 := ListNode{Val: 4, Next: &x1}
	// x3 := ListNode{Val: 2, Next: &x2}
	// y1 := ListNode{Val: 4, Next: nil}
	// y2 := ListNode{Val: 6, Next: &y1}
	// y3 := ListNode{Val: 5, Next: &y2}
	// addTwoNumbers1(&x3, &y3).printList()

	x1 := ListNode{Val: 1, Next: nil}
	y1 := ListNode{Val: 9, Next: nil}
	y2 := ListNode{Val: 9, Next: &y1}
	addTwoNumbers1(&x1, &y2).printList()
}
