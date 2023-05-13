/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/722/
*/

package task0026

func Solution(nums []int) int {
	length := len(nums)
	total := (length * (length + 1)) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return total - sum
}
