package task0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0005Solution(t *testing.T) {
	tests := []struct {
		name   string
		number int
		whant  int
	}{
		{
			name:   "Multiples of 3 and 5 for 10",
			number: 10,
			whant:  23,
		},
		{
			name:   "Multiples of 3 and 5 for -10",
			number: -10,
			whant:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.number)
			assert.Equal(t, res, tt.whant)
		})
	}
}
