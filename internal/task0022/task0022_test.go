package task0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0022Solution(t *testing.T) {
	tests := []struct {
		name  string
		start string
		end   string
		want  int
	}{
		{
			name:  "example 1",
			start: "10.0.0.0",
			end:   "10.0.0.50",
			want:  50,
		},
		{
			name:  "example 2",
			start: "20.0.0.10",
			end:   "20.0.1.0",
			want:  246,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.start, tt.end)
			assert.Equal(t, res, tt.want)
		})
	}
}
