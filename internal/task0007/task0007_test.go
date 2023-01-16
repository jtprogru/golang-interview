package task0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testHeadForNode5 = makeFirstTestHead()
	testHeadForNode1 = makeSecondTestHead()
)

func TestTask0006Solution(t *testing.T) {
	tests := []struct {
		name  string
		head  *ListNode
		node  *ListNode
		whant *ListNode
	}{
		{
			name: "head = [4,5,1,9], node = 5",
			head: testHeadForNode5,
			node: testHeadForNode5.Next,
			whant: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
		{
			name: "head = [4,5,1,9], node = 1",
			head: testHeadForNode1,
			node: testHeadForNode1.Next.Next,
			whant: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solution(tt.node)
			assert.Equal(t, tt.head, tt.whant)
		})
	}
}

func makeFirstTestHead() *ListNode {
	return &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
}

func makeSecondTestHead() *ListNode {
	return &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

}
