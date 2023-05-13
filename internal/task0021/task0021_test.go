package task0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0021Solution(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       []int
	}{
		{
			name:       "example 1",
			expression: "2-1-1",
			want:       []int{0, 2},
		},
		{
			name:       "example 2",
			expression: "2*3-4*5",
			want:       []int{-34, -14, -10, -10, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.expression)
			assert.Equal(t, res, tt.want)
		})
	}
}
