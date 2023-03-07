package task0016

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0016Solution(t *testing.T) {
	tests := []struct {
		name      string
		morseCode string
		whant     string
	}{
		{
			name:      "example 1",
			morseCode: ".... . -.--   .--- ..- -.. .",
			whant:     "HEY JUDE",
		},
		{
			name:      "example 2",
			morseCode: "... --- ... -.-.--",
			whant:     "SOS!",
		},
		{
			name:      "example 3",
			morseCode: "... --- ... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-",
			whant:     "SOS! THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Solution(tt.morseCode)
			fmt.Println(res)
			assert.Equal(t, res, tt.whant)
		})
	}
}
