package task0016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0016Solution(t *testing.T) {
	tests := []struct {
		name      string
		morseCode string
		want      string
	}{
		{
			name:      "example 1",
			morseCode: ".... . -.--   .--- ..- -.. .",
			want:      "HEY JUDE",
		},
		{
			name:      "example 2",
			morseCode: "... --- ... -.-.--",
			want:      "SOS!",
		},
		{
			name:      "example 3",
			morseCode: "... --- ... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-",
			want:      "SOS! THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.morseCode)
			assert.Equal(t, res, tt.want)
		})
	}
}
