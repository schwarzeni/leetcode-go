package _0042

// 2020.05.01
// https://leetcode-cn.com/problems/trapping-rain-water/
// 没写出来

//使用双指针
func trap(height []int) int {
	size, li, ri, lmax, rmax := 0, 0, len(height)-1, -1, -1

	for li < ri {
		if height[li] < height[ri] {
			if lmax == -1 || height[li] > lmax {
				lmax = height[li]
			} else {
				size += lmax - height[li]
			}
			li++
		} else {
			if rmax == -1 || height[ri] > rmax {
				rmax = height[ri]
			} else {
				size += rmax - height[ri]
			}
			ri--
		}
	}

	return size
}

// 使用栈
func trap2(height []int) int {
	size := 0
	stack := []int{}

	for idx, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := idx - stack[len(stack)-1] - 1
			boundedHeight := min(v, height[stack[len(stack)-1]]) - height[top]
			size += distance * boundedHeight
		}
		stack = append(stack, idx)
	}

	return size
}

// 动态规划
func trap3(height []int) int {
	size, n := 0, len(height)
	maxls, maxrs := make([]int, n), make([]int, n)
	if n == 0 {
		return 0
	}
	maxls[0] = height[0]
	for i := 1; i < n; i++ {
		maxls[i] = max(maxls[i-1], height[i])
	}
	maxrs[n-1] = height[n-1]
	for j := n - 2; j >= 0; j-- {
		maxrs[j] = max(maxrs[j+1], height[j])
	}

	for idx, v := range height {
		if idx == 0 || idx == n-1 {
			continue
		}
		size += min(maxls[idx], maxrs[idx]) - v
	}

	return size
}

func max(d1, d2 int) int {
	if d1 > d2 {
		return d1
	}
	return d2
}

func min(d1, d2 int) int {
	if d1 > d2 {
		return d2
	}
	return d1
}
