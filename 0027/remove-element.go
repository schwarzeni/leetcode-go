package _0027

// 2020.04.26
// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var idx int
	for _, v := range nums {
		if v != val {
			nums[idx] = v
			idx++
		}
	}
	return idx
}
