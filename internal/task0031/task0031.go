package task0031

// Функция для сортировки выбором
func SelectionSort(arr *[]int) {
	if arr == nil || len(*arr) <= 1 {
		return
	}

	n := len(*arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		// Обмен значениями
		(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
	}
}
