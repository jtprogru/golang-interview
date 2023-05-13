/*
https://leetcode.com/problems/min-stack/description/
*/

package task0030

import "math"

type MinStack struct {
	stk1 []int
	stk2 []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt32}}
}

func (ms *MinStack) Push(val int) {
	ms.stk1 = append(ms.stk1, val)
	ms.stk2 = append(ms.stk2, min(val, ms.stk2[len(ms.stk2)-1]))
}

func (ms *MinStack) Pop() {
	ms.stk1 = ms.stk1[:len(ms.stk1)-1]
	ms.stk2 = ms.stk2[:len(ms.stk2)-1]
}

func (ms *MinStack) Top() int {
	return ms.stk1[len(ms.stk1)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.stk2[len(ms.stk2)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
