package task0012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0012Solution(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        [][]int
	}{
		{
			name:        "example 1",
			n:           4,
			connections: [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
			want:        [][]int{{1, 3}},
		},
		{
			name:        "example 2",
			n:           2,
			connections: [][]int{{0, 1}},
			want:        [][]int{{0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.n, tt.connections)
			assert.Equal(t, res, tt.want)
		})
	}
}
