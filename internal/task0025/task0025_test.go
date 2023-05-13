package task0025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0025Solution(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "example 1",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "example 2",
			numRows: 1,
			want: [][]int{
				{1},
			},
		},
		{
			name:    "example 3",
			numRows: 10,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
				{1, 8, 28, 56, 70, 56, 28, 8, 1},
				{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.numRows)
			assert.Equal(t, res, tt.want)
		})
	}
}
