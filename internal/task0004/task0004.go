/*
https://www.codewars.com/kata/55f2b110f61eb01779000053/train/go
*/

package task0004

func Solution(a, b int) int {
	if a == b {
		return a
	}

	if a > b {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}
