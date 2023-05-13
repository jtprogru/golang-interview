package task0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0026Solution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.nums)
			assert.Equal(t, res, tt.want)
		})
	}
}
