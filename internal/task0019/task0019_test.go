package task0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0019Solution(t *testing.T) {
	tests := []struct {
		name  string
		strng string
		want  string
	}{
		{
			name:  "example 1",
			strng: "103 123 4444 99 2000",
			want:  "2000 103 123 4444 99",
		},
		{
			name:  "example 2",
			strng: "2000 10003 1234000 44444444 9999 11 11 22 123",
			want:  "11 11 2000 10003 22 123 1234000 44444444 9999",
		},
		{
			name:  "example 3",
			strng: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.strng)
			assert.Equal(t, res, tt.want)
		})
	}
}
