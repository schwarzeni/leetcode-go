package _0026

// 2020.04.26
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	var i, curr int

	if len(nums) < 2 {
		return len(nums)
	}
	curr = nums[0]
	for _, v := range nums {
		if v != curr {
			i++
			nums[i], curr = v, v
		}
	}

	return i + 1
}
