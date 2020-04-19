package _0056

import "sort"

// 2020.04.19
// https://leetcode-cn.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	var (
		result [][]int
		l, r   int
	)
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	l = intervals[0][0]
	r = intervals[0][1]

	for idx, interval := range intervals {
		cl, cr := interval[0], interval[1]
		if cl > r {
			result = append(result, []int{l, r})
			l, r = cl, cr
		} else if cr >= r {
			r = cr
		}
		if idx == len(intervals)-1 {
			result = append(result, []int{l, r})
		}
	}

	return result
}
