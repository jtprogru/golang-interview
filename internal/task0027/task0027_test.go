package task0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0027Solution(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		zero int
		one  int
		want int
	}{
		{
			name: "example 1",
			low:  3,
			high: 3,
			zero: 1,
			one:  1,
			want: 8,
		},
		{
			name: "example 2",
			low:  2,
			high: 3,
			zero: 1,
			one:  2,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.low, tt.high, tt.zero, tt.one)
			assert.Equal(t, res, tt.want)
		})
	}
}
