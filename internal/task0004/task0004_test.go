package task0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0004Solution(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "Get sum between 0 and 1",
			a:    0,
			b:    1,
			want: 1,
		},
		{
			name: "Get sum between 1 and 2",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "Get sum between 5 and -1",
			a:    5,
			b:    -1,
			want: 14,
		},
		{
			name: "Get sum between 505 and 4",
			a:    505,
			b:    4,
			want: 127759,
		},
		{
			name: "Get sum between 321 and 123",
			a:    321,
			b:    123,
			want: 44178,
		},
		{
			name: "Get sum between 0 and -1",
			a:    0,
			b:    -1,
			want: -1,
		},
		{
			name: "Get sum between -50 and 0",
			a:    -50,
			b:    0,
			want: -1275,
		},
		{
			name: "Get sum between -1 and -5",
			a:    -1,
			b:    -5,
			want: -15,
		},
		{
			name: "Get sum between -5 and -5",
			a:    -5,
			b:    -5,
			want: -5,
		},
		{
			name: "Get sum between -505 and 4",
			a:    -505,
			b:    4,
			want: -127755,
		},
		{
			name: "Get sum between -321 and 123",
			a:    -321,
			b:    123,
			want: -44055,
		},
		{
			name: "Get sum between 0 and 0",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "Get sum between -5 and -1",
			a:    -5,
			b:    -1,
			want: -15,
		},
		{
			name: "Get sum between 5 and 1",
			a:    5,
			b:    1,
			want: 15,
		},
		{
			name: "Get sum between -17 and -17",
			a:    -17,
			b:    -17,
			want: -17,
		},
		{
			name: "Get sum between 17 and 17",
			a:    17,
			b:    17,
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.a, tt.b)
			assert.Equal(t, res, tt.want)
		})
	}
}
