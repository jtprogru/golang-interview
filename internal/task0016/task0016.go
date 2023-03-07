package task0016

import "strings"

var morsemap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
}

var characterMap = map[string]string{}

func decodeLetter(code string) string {
	return characterMap[code]
}

func decodeMessage(message, letterSplitter string) string {
	var decodedMessage string
	var letters = strings.Split(message, letterSplitter)
	for _, letter := range letters {
		decodedMessage += decodeLetter(strings.Trim(letter, " "))
	}
	return decodedMessage
}

func Solution(morseCode string) string {
	var res string
	for key, value := range morsemap {
		characterMap[value] = key
	}
	words := strings.Split(morseCode, "   ")
	for _, letter := range words {
		res += decodeMessage(letter, " ") + " "
	}
	return strings.Trim(res, " ")
}
