package task0017

import (
	"sort"
)

func Solution(s1 string, s2 string) string {
	uniqueChars := [26]bool{}

	for _, c := range s1 {
		uniqueChars[c-'a'] = true
	}

	for _, c := range s2 {
		uniqueChars[c-'a'] = true
	}

	var result []rune
	for i, exists := range uniqueChars {
		if exists {
			result = append(result, rune(i)+'a')
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return string(result)
}
