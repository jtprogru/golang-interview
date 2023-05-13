package task0030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0030Solution(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "example 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minStack := Constructor()
			minStack.Push(-2)
			minStack.Push(0)
			minStack.Push(-3)
			assert.Equal(t, minStack.GetMin(), -3) // return -3
			minStack.Pop()
			assert.Equal(t, minStack.Top(), 0)     // return 0
			assert.Equal(t, minStack.GetMin(), -2) // return -2
		})
	}
}
