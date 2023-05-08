/*
https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/603/
*/

package task0008

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}
