package task0017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0017Solution(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want string
	}{
		{
			name: "example 1",
			s1:   "aretheyhere",
			s2:   "yestheyarehere",
			want: "aehrsty",
		},
		{
			name: "example 2",
			s1:   "loopingisfunbutdangerous",
			s2:   "lessdangerousthancoding",
			want: "abcdefghilnoprstu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.s1, tt.s2)
			assert.Equal(t, res, tt.want)
		})
	}
}
