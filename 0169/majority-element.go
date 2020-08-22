package _0169

// 2020.08.22
// https://leetcode-cn.com/problems/majority-element/
func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for i := range nums {
		if cnt == 0 {
			res = nums[i]
			cnt++
		} else if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
