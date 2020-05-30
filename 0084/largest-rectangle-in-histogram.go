package _0084

// 2020.05.30
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	max := 0
	for idx, h := range heights {
		w := 1
		for i := idx - 1; i >= 0 && heights[i] >= h; i-- {
			w++
		}
		for i := idx + 1; i < len(heights) && heights[i] >= h; i++ {
			w++
		}
		if s := w * h; s > max {
			max = s
		}
	}
	return max
}
