package task0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0001Solution(t *testing.T) {
	tests := []struct {
		name   string
		person *Person
		want   Resp
	}{
		{
			name:   "First",
			person: &Person{Name: "Bob"},
			want: Resp{
				First: &Person{
					Name: "Alice",
				},
				Second: &Person{
					Name: "John",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.person)
			assert.Equal(t, res, tt.want)
		})
	}
}
