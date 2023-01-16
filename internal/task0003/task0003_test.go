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
		whant  bool
	}{
		{
			name:   "abc ending bc",
			str:    "abc",
			ending: "bc",
			whant:  true,
		},
		{
			name:   "abc not ending d",
			str:    "abc",
			ending: "d",
			whant:  false,
		},
		{
			name:   "empty string ending empty string",
			str:    "",
			ending: "",
			whant:  true,
		},
		{
			name:   "single space ending empty string",
			str:    " ",
			ending: "",
			whant:  true,
		},
		{
			name:   "abc ending c",
			str:    "abc",
			ending: "c",
			whant:  true,
		},
		{
			name:   "banana ending ana",
			str:    "banana",
			ending: "ana",
			whant:  true,
		},
		{
			name:   "a not ending z",
			str:    "a",
			ending: "z",
			whant:  false,
		},
		{
			name:   "empty string not ending t",
			str:    "",
			ending: "t",
			whant:  false,
		},
		{
			name:   "$a = $b + 1 not ending +1",
			str:    "$a = $b + 1",
			ending: "+1",
			whant:  false,
		},
		{
			name:   "some spaces ending some spaces",
			str:    "    ",
			ending: "   ",
			whant:  true,
		},
		{
			name:   "1oo not ending 100",
			str:    "1oo",
			ending: "100",
			whant:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.str, tt.ending)
			assert.Equal(t, res, tt.whant)
		})
	}
}
