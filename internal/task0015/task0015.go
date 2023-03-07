// Create a function taking a positive integer between 1 and 3999 (both included)
// as its parameter and returning a string containing the Roman Numeral
// representation of that integer.
// Modern Roman numerals are written by expressing each digit separately starting
// with the left most digit and skipping any digit with a value of zero.
// In Roman numerals 1990 is rendered: 1000=M, 900=CM, 90=XC; resulting in MCMXC.
// 2008 is written as 2000=MM, 8=VIII; or MMVIII. 1666 uses each Roman symbol
// in descending order: MDCLXVI.
// Example:
// solution(1000); // should return 'M'
// Help:
// Symbol    Value
// I          1
// V          5
// X          10
// L          50
// C          100
// D          500
// M          1,000
// Remember that there can't be more than 3 identical symbols in a row.
// More about roman numerals - http://en.wikipedia.org/wiki/Roman_numerals

package task0015

import (
	"fmt"
	"math"
	"strings"
)

var (
	ARABIC_NUMBERS = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ROMAN_NUMBERS  = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// A tuple (pair) holding our numeric value and its roman counterpart.
type NumeralTuple struct {
	arabic int
	roman  string
}

func Solution(number int) string {
	if number < 0 {
		return ""
	}

	// See: https://en.wikipedia.org/wiki/Roman_numerals#Zero
	if number == 0 {
		return "nulla"
	}

	result := []string{}
	numerals, _ := zip(ARABIC_NUMBERS, ROMAN_NUMBERS)

	var quotient int

	for _, numeral := range numerals {
		quotient, number = divmod(number, numeral.arabic)

		// strings.Repeat() takes care of zeroes so we don't have to add
		// conditional clauses.
		result = append(result, strings.Repeat(numeral.roman, quotient))
	}

	// Join the result slice to string eliminating all empty values, and
	// thus returning the final desired value.
	return strings.Join(result, "")
}

func zip(a []int, b []string) ([]NumeralTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("given lists have different length")
	}

	result := make([]NumeralTuple, len(a))

	for i, e := range a {
		result[i] = NumeralTuple{e, b[i]}
	}

	return result, nil
}

// Returns (a / b, a % b) as a tuple.
func divmod(a int, b int) (int, int) {
	quotient := float64(a / b)
	remainder := math.Mod(float64(a), float64(b))

	return int(quotient), int(remainder)
}
