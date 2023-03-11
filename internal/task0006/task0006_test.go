package task0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0006Solution(t *testing.T) {
	tests := []struct {
		name string
		prod uint64
		want Resp
	}{
		{
			name: "Sum Fib for 4895",
			prod: 4895,
			want: Resp{55, 89, 1},
		},
		{
			name: "Sum Fib for 5895",
			prod: 5895,
			want: Resp{89, 144, 0},
		},
		{
			name: "Sum Fib for 74049690",
			prod: 74049690,
			want: Resp{6765, 10946, 1},
		},
		{
			name: "Sum Fib for 84049690",
			prod: 84049690,
			want: Resp{10946, 17711, 0},
		},
		{
			name: "Sum Fib for 714",
			prod: 714,
			want: Resp{21, 34, 1},
		},
		{
			name: "Sum Fib for 800",
			prod: 800,
			want: Resp{34, 55, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.prod)
			assert.Equal(t, res, tt.want)
		})
	}
}
