package task0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0008Solution(t *testing.T) {
	tests := []struct {
		name  string
		head  *ListNode
		n     int
		whant *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], n = 2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			n: 2,
			whant: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "head = [1], n = 1",
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			n:     1,
			whant: nil,
		},
		{
			name: "head = [1,2], n = 1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			n: 1,
			whant: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.head, tt.n)
			assert.Equal(t, res, tt.whant)
		})
	}
}
