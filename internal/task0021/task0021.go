/*
https://leetcode.com/problems/different-ways-to-add-parentheses/description/
*/

package task0021

import (
	"sort"
	"strconv"
)

func Solution(expression string) []int {
	return compute(expression)
}

func compute(expression string) []int {
	result := make([]int, 0)

	for i := 0; i < len(expression); i++ {
		c := expression[i]
		if c == '+' || c == '-' || c == '*' {
			left := compute(expression[0:i])
			right := compute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					switch c {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}
	sort.Ints(result)
	return result
}
