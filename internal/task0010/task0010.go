/*
https://leetcode.com/problems/insert-interval/
*/

package task0010

type Interval struct {
	Start int
	End   int
}

func Solution(intervals [][]int, newInterval []int) [][]int {
	ilist := []Interval{}
	for _, v := range intervals {
		tmp := Interval{v[0], v[1]}
		ilist = append(ilist, tmp)
	}
	newItv := Interval{newInterval[0], newInterval[1]}
	resList := insertIntervals(ilist, newItv)
	res := [][]int{}
	for _, v := range resList {
		tmp := []int{v.Start, v.End}
		res = append(res, tmp)
	}
	return res
}

func insertIntervals(a []Interval, newInterval Interval) []Interval {
	res := []Interval{}
	if len(a) == 0 {
		res = append(res, newInterval)
		return res
	}
	cur := 0
	for cur < len(a) && a[cur].End < newInterval.Start {
		res = append(res, a[cur])
		cur++
	}
	for cur < len(a) && a[cur].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, a[cur].Start), End: max(newInterval.End, a[cur].End)}
		cur++
	}
	res = append(res, newInterval)
	for cur < len(a) {
		res = append(res, a[cur])
		cur++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
