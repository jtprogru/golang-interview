/*
You are given two non-empty linked lists representing
two non-negative integers. The digits are stored in
reverse order, and each of their nodes contains a single
digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading
zero, except the number 0 itself.
*/

package task0014

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
