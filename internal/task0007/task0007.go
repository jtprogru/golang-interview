/*
https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/553/
*/

package task0007

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(node *ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}
