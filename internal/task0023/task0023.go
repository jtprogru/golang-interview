/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/560/
*/

package task0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := Solution(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
