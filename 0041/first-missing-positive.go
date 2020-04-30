package _0041

// 2020.04.30
// https://leetcode-cn.com/problems/first-missing-positive/
// 先将输入数组中小于1的数全设为1，再用这个数组数据的正负来计数，这样就不用额外的空间了，nbnb，学到了
func firstMissingPositive(nums []int) int {
	// 查找大于0的最大值和最小值
	max, min := -1, -1
	for _, v := range nums {
		if v > 0 {
			if max == -1 || v > max {
				max = v
			}
			if min == -1 || v < min {
				min = v
			}
		}
	}

	// 1 [4,5,10,17] 或 [-1, 0]
	if min > 1 || max < 1 {
		return 1
	}

	// 2. [-3, -1, 1, 4, 7]
	// 清理脏数据
	for i, v := range nums {
		if v <= 0 {
			nums[i] = 1
		}
	}

	// 使用自身来计数
	for _, v := range nums {
		if abs(v) <= len(nums) {
			if nums[abs(v)-1] > 0 {
				nums[abs(v)-1] *= -1
			}
		}
	}

	// 检测数据
	for idx, v := range nums {
		if v > 0 {
			return idx + 1
		}
	}

	// 3. [1 2 3 4 5 6 7]
	return len(nums) + 1
}

func abs(d int) int {
	if d < 0 {
		return -d
	}
	return d
}
