package _0057

// 2020.05.04
// https://leetcode-cn.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) [][]int {
	l, r, result := -1, -1, [][]int{}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	for idx, interval := range intervals {
		if idx == 0 && interval[0] > newInterval[1] {
			l, r = newInterval[0], newInterval[1]
			result = append(result, newInterval)
		}
		if interval[1] < newInterval[0] || interval[0] > newInterval[1] {
			result = append(result, interval)
		}
		if interval[1] < newInterval[0] && (idx == len(intervals)-1 || intervals[idx+1][0] > newInterval[1]) {
			l, r = newInterval[0], newInterval[1]
			result = append(result, newInterval)
		}
		if l == -1 {
			if (idx == 0 || newInterval[0] > intervals[idx-1][1]) && newInterval[0] < interval[0] {
				l = newInterval[0]
			} else if newInterval[0] >= interval[0] && newInterval[0] <= interval[1] {
				l = interval[0]
			}
		}
		if r == -1 {
			if (idx == len(intervals)-1 || intervals[idx+1][0] > newInterval[1]) && newInterval[1] > interval[1] {
				r = newInterval[1]
				result = append(result, []int{l, r})
			} else if interval[1] >= newInterval[1] && interval[0] <= newInterval[1] {
				r = interval[1]
				result = append(result, []int{l, r})
			}
		}
	}

	return result
}
