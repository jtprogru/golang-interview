package task0010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0010Solution(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.intervals, tt.newInterval)
			assert.Equal(t, res, tt.want)
		})
	}
}
