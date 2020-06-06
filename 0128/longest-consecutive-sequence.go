package _0128

// https://leetcode-cn.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	mark := make(map[int]bool)
	for i, _ := range nums {
		mark[nums[i]] = false
	}
	maxLen := 0

	for i, _ := range nums {
		if mark[nums[i]] == true {
			continue
		}
		mark[nums[i]] = true
		l, v := 1, nums[i]
		for {
			v--
			visited, ok := mark[v]
			if !ok || visited {
				break
			}
			l++
			if ok {
				mark[v] = true
			}
		}
		v = nums[i]
		for {
			v++
			visited, ok := mark[v]
			if !ok || visited {
				break
			}
			l++
			if ok {
				mark[v] = true
			}
		}
		if maxLen < l {
			maxLen = l
		}
	}

	return maxLen
}
