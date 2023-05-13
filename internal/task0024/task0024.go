/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/771/
*/

package task0024

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(list1 *ListNode, list2 *ListNode) *ListNode {
	// Создадим фиктивный узел
	dummyNode := &ListNode{}
	current := dummyNode

	// Проходим по двум спискам
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Если один из списков еще содержит узлы, добавляем их в конец результата
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// Возвращаем объединенный список, исключая фиктивный узел
	return dummyNode.Next
}
