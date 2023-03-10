package task0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0011Solution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "example 1",
			s:    "abcdefg",
			k:    2,
			want: "bacdfeg",
		},
		{
			name: "example 2",
			s:    "abcd",
			k:    2,
			want: "bacd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.s, tt.k)
			assert.Equal(t, res, tt.want)
		})
	}
}
