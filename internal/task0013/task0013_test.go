package task0013

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0013Solution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		whant  []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			whant:  []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			whant:  []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			whant:  []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.nums, tt.target)
			fmt.Printf("result is: %+v", res)
			assert.Equal(t, res, tt.whant)
		})
	}
}
