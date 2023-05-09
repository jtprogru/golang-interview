/*
https://leetcode.com/problems/search-a-2d-matrix-ii/description/
*/

package task0020

func Solution(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	if matrix == nil || m == 0 || n == 0 {
		return false
	}
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
