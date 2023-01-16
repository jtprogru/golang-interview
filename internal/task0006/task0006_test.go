package task0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0005Solution(t *testing.T) {
	tests := []struct {
		name  string
		prod  uint64
		whant Resp
	}{
		{
			name:  "Sum Fib for 4895",
			prod:  4895,
			whant: Resp{55, 89, 1},
		},
		{
			name:  "Sum Fib for 5895",
			prod:  5895,
			whant: Resp{89, 144, 0},
		},
		{
			name:  "Sum Fib for 74049690",
			prod:  74049690,
			whant: Resp{6765, 10946, 1},
		},
		{
			name:  "Sum Fib for 84049690",
			prod:  84049690,
			whant: Resp{10946, 17711, 0},
		},
		{
			name:  "Sum Fib for 714",
			prod:  714,
			whant: Resp{21, 34, 1},
		},
		{
			name:  "Sum Fib for 800",
			prod:  800,
			whant: Resp{34, 55, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.prod)
			assert.Equal(t, res, tt.whant)
		})
	}
}
