package task0018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0018Solution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    12,
			want: 21,
		},
		{
			name: "example 2",
			n:    513,
			want: 531,
		},
		{
			name: "example 3",
			n:    2017,
			want: 2071,
		},
		{
			name: "example 4",
			n:    414,
			want: 441,
		},
		{
			name: "example 5",
			n:    144,
			want: 414,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.n)
			assert.Equal(t, res, tt.want)
		})
	}
}
