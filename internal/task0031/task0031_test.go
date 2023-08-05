package task0031

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{nil, nil},
	}

	for _, tc := range testCases {
		arr := tc.input
		SelectionSort(&arr)
		if !reflect.DeepEqual(arr, tc.expected) {
			t.Errorf("Ошибка в тесте с массивом %v. Ожидалось %v, получено %v", tc.input, tc.expected, arr)
		}
	}
}
