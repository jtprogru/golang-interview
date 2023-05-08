package task0019

import (
	"sort"
	"strconv"
	"strings"
)

func Solution(strng string) string {
	weights := strings.Fields(strng)
	sort.SliceStable(weights, func(i, j int) bool {
		wi, wj := getWeight(weights[i]), getWeight(weights[j])
		if wi == wj {
			return weights[i] < weights[j]
		}
		return wi < wj
	})
	return strings.Join(weights, " ")
}

func getWeight(s string) int {
	weight := 0
	for _, c := range s {
		digit, _ := strconv.Atoi(string(c))
		weight += digit
	}
	return weight
}
