package task0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0009Solution(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		whant bool
	}{
		{
			name:  "x = 121",
			x:     121,
			whant: true,
		},
		{
			name:  "x = -121",
			x:     -121,
			whant: false,
		},
		{
			name:  "x = 10",
			x:     10,
			whant: false,
		},
		{
			name:  "x = 99",
			x:     99,
			whant: true,
		},
		{
			name:  "x = 987789",
			x:     987789,
			whant: true,
		},
		{
			name:  "x = 9876789",
			x:     9876789,
			whant: true,
		},
		{
			name:  "x = 1",
			x:     1,
			whant: true,
		},
		{
			name:  "x = 12322",
			x:     12322,
			whant: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.x)
			assert.Equal(t, res, tt.whant)
		})
	}
}
