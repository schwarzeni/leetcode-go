package _5497

// 2020.08.23
// https://leetcode-cn.com/problems/find-latest-group-of-size-m/
func findLatestStep(arr []int, m int) int {
	// 当前 idx 所在1列表最左端的 idx ，如果不存在则为-1
	larr := make([]int, len(arr)+3)
	// 当前 idx 所在1列表最右端的 idx ，如果不存在则为-1
	rarr := make([]int, len(arr)+3)
	for i := range larr {
		larr[i], rarr[i] = -1, -1
	}
	count := 0
	lastStep := -1

	for i, idx := range arr {
		switch {
		case larr[idx-1] == -1 && larr[idx+1] == -1:
			larr[idx], rarr[idx] = idx, idx
			if m == 1 {
				count++
			}
		case larr[idx-1] != -1 && larr[idx+1] == -1:
			rarr[larr[idx-1]] = idx
			rarr[idx] = idx
			larr[idx] = larr[idx-1]
			if idx-larr[idx]+1 == m {
				count++
			}
			if idx-larr[idx] == m {
				count--
			}
		case larr[idx-1] == -1 && larr[idx+1] != -1:
			larr[rarr[idx+1]] = idx
			larr[idx] = idx
			rarr[idx] = rarr[idx+1]
			if rarr[idx]-idx+1 == m {
				count++
			}
			if rarr[idx]-idx == m {
				count--
			}
		case larr[idx-1] != -1 && larr[idx+1] != -1:
			if idx-1-larr[idx-1]+1 == m {
				count--
			}
			if rarr[idx+1]-(idx+1)+1 == m {
				count--
			}
			larr[rarr[idx+1]] = larr[idx-1]
			rarr[larr[idx-1]] = rarr[idx+1]
			if rarr[idx+1]-larr[idx-1]+1 == m {
				count++
			}
		}
		if count > 0 {
			lastStep = i + 1
		}
	}
	return lastStep
}
