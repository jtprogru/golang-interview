package task0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0003Solution(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		ending string
		want   bool
	}{
		{
			name:   "abc ending bc",
			str:    "abc",
			ending: "bc",
			want:   true,
		},
		{
			name:   "abc not ending d",
			str:    "abc",
			ending: "d",
			want:   false,
		},
		{
			name:   "empty string ending empty string",
			str:    "",
			ending: "",
			want:   true,
		},
		{
			name:   "single space ending empty string",
			str:    " ",
			ending: "",
			want:   true,
		},
		{
			name:   "abc ending c",
			str:    "abc",
			ending: "c",
			want:   true,
		},
		{
			name:   "banana ending ana",
			str:    "banana",
			ending: "ana",
			want:   true,
		},
		{
			name:   "a not ending z",
			str:    "a",
			ending: "z",
			want:   false,
		},
		{
			name:   "empty string not ending t",
			str:    "",
			ending: "t",
			want:   false,
		},
		{
			name:   "$a = $b + 1 not ending +1",
			str:    "$a = $b + 1",
			ending: "+1",
			want:   false,
		},
		{
			name:   "some spaces ending some spaces",
			str:    "    ",
			ending: "   ",
			want:   true,
		},
		{
			name:   "1oo not ending 100",
			str:    "1oo",
			ending: "100",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.str, tt.ending)
			assert.Equal(t, res, tt.want)
		})
	}
}
