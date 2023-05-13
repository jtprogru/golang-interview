package task0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0028Solution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example 3",
			s:    "pwwkew",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.s)
			assert.Equal(t, res, tt.want)
		})
	}
}
