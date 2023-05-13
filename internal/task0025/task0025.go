/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/601/
*/

package task0025

func Solution(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// Создаем новый массив для каждой строки
		triangle[i] = make([]int, i+1)
		// Первый и последний элемент каждой строки равен 1
		triangle[i][0] = 1
		triangle[i][i] = 1

		// Заполняем остальные элементы строки
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}
