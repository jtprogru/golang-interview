package task0029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0029Solution(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name: "example 1",
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "example 2",
			intervals: [][]int{
				{1, 4},
				{4, 5},
			},
			want: [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.intervals)
			assert.Equal(t, res, tt.want)
		})
	}
}
