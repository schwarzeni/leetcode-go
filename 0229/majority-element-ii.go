package _0229

// 2020.09.17
// https://leetcode-cn.com/problems/majority-element-ii/
func majorityElement(nums []int) []int {
	var res []int
	if len(nums) < 1 {
		return res
	}

	c1, c2, m1, m2 := 0, 0, nums[0], nums[0]
	for i := range nums {
		num := nums[i]
		switch {
		case num == m1:
			c1++
		case num == m2:
			c2++
		case c1 == 0:
			c1, m1 = 1, num
		case c2 == 0:
			c2, m2 = 1, num
		default:
			c1, c2 = c1-1, c2-1
		}
	}

	c1, c2 = 0, 0
	for i := range nums {
		switch nums[i] {
		case m1:
			c1++
		case m2:
			c2++
		}
	}
	if c1 > len(nums)/3 {
		res = append(res, m1)
	}
	if c2 > len(nums)/3 && m1 != m2 {
		res = append(res, m2)
	}
	return res
}
