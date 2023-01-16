package task0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0002Solution(t *testing.T) {
	tests := []struct {
		name  string
		start int
		end   int
		whant *Resp
	}{
		{
			name:  "First",
			start: 0,
			end:   10,
			whant: &Resp{
				A: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name:  "Second",
			start: 0,
			end:   2,
			whant: &Resp{
				A: []int{0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.start, tt.end)
			assert.Equal(t, res, tt.whant)
		})
	}
}
