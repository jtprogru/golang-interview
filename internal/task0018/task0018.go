package task0018

import (
	"sort"
	"strconv"
)

func Solution(n int) int {
	digits := getDigits(n)
	length := len(digits)

	for i := length - 1; i > 0; i-- {
		if digits[i-1] < digits[i] {
			for j := length - 1; j >= i; j-- {
				if digits[j] > digits[i-1] {
					digits[i-1], digits[j] = digits[j], digits[i-1]
					sort.Ints(digits[i:])
					return digitsToInt(digits)
				}
			}
		}
	}

	return -1
}

func getDigits(n int) []int {
	str := strconv.Itoa(n)
	digits := make([]int, len(str))

	for i, c := range str {
		d, _ := strconv.Atoi(string(c))
		digits[i] = d
	}

	return digits
}

func digitsToInt(digits []int) int {
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	return result
}
