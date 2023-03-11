package task0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0009Solution(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "x = 121",
			x:    121,
			want: true,
		},
		{
			name: "x = -121",
			x:    -121,
			want: false,
		},
		{
			name: "x = 10",
			x:    10,
			want: false,
		},
		{
			name: "x = 99",
			x:    99,
			want: true,
		},
		{
			name: "x = 987789",
			x:    987789,
			want: true,
		},
		{
			name: "x = 9876789",
			x:    9876789,
			want: true,
		},
		{
			name: "x = 1",
			x:    1,
			want: true,
		},
		{
			name: "x = 12322",
			x:    12322,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.x)
			assert.Equal(t, res, tt.want)
		})
	}
}
