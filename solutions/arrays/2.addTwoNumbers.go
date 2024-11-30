package arrays

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) printList() {
	for cur := n; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			fmt.Printf("%d \n", cur.Val)
		} else {
			fmt.Printf("%d -> ", cur.Val)
		}
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 *
 * KEYWORDS: single digits, reverse order
 * treat one list as a single number, think how you would add up two numbers
 */

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var curNodePtr *ListNode
	result := &ListNode{Val: 0, Next: nil}
	p1 := l1
	p2 := l2
	var carry int
	curNodePtr = result
	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			curNodePtr.Val = (p1.Val + p2.Val + carry) % 10
			carry = (p1.Val + p2.Val + carry) / 10
			if p1.Next == nil && p2.Next == nil {
				break
			}
			p1 = p1.Next
			p2 = p2.Next
		} else if p1 != nil && p2 == nil {
			// x1 -> x2 -> x3
			// y1 -> y2
			curNodePtr.Val = (p1.Val + carry) % 10
			carry = (p1.Val + carry) / 10
			if p1.Next == nil {
				break
			}
			p1 = p1.Next
		} else if p1 == nil && p2 != nil {
			// y1 -> y2 -> y3
			// x1 -> x2
			curNodePtr.Val = (p2.Val + carry) % 10
			carry = (p2.Val + carry) / 10
			if p2.Next == nil {
				break
			}
			p2 = p2.Next
		}

		curNodePtr.Next = &ListNode{Val: carry, Next: nil}
		curNodePtr = curNodePtr.Next
	}
	if carry > 0 {
		curNodePtr.Next = &ListNode{Val: carry, Next: nil}
	}

	return result
}
