package _0046

// 2020.05.02
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	// idxMap: numIdx --> resultIdx+1
	var f func(level int, idxMap []int)
	var result [][]int

	f = func(level int, idxMap []int) {
		if level == len(nums) {
			r := make([]int, len(nums))
			for i, v := range idxMap {
				r[v-1] = nums[i]
			}
			result = append(result, r)
			return
		}
		for i := 0; i < len(nums); i++ {
			if idxMap[i] == 0 {
				idxMap[i] = level + 1
				f(level+1, idxMap)
				idxMap[i] = 0
			}
		}
	}

	f(0, make([]int, len(nums)))

	return result
}
