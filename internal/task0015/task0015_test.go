package task0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0015Solution(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want string
	}{
		{
			name: "example 1",
			x:    182,
			want: "CLXXXII",
		},
		{
			name: "example 2",
			x:    1990,
			want: "MCMXC",
		},
		{
			name: "example 3",
			x:    1875,
			want: "MDCCCLXXV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.x)
			assert.Equal(t, res, tt.want)
		})
	}
}
