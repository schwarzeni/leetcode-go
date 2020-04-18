package _0011

// 2020.04.18
// https://leetcode-cn.com/problems/container-with-most-water/
// O(n^2)
//func maxArea(height []int) int {
//	max := 0
//	for i := len(height) - 2; i >= 0; i-- {
//		for j := i + 1; j < len(height); j++ {
//			c := 0
//			if height[i] > height[j] {
//				c = height[j] * (j - i)
//			} else {
//				c = height[i] * (j - i)
//			}
//			if c > max {
//				max = c
//			}
//		}
//	}
//	return max
//}

// 双指针法
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1

	for i < j {
		r := 0
		if height[i] > height[j] {
			r = height[j] * (j - i)
			j--
		} else {
			r = height[i] * (j - i)
			i++
		}
		if r > max {
			max = r
		}
	}
	return max
}
