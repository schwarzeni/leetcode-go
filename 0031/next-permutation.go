package _31

// 2020.04.06
// 注意关注旋转数组的操作
// https://leetcode-cn.com/problems/next-permutation/
func nextPermutation(nums []int) {
	var (
		size = len(nums)
		idx  int
	)
	if size == 1 {
		return
	}

	// 第一次从后向前遍历，找到一个数，这个数比后一个数小，记为 k
	for idx = size - 2; idx >= 0; idx-- {
		if nums[idx] < nums[idx+1] {
			break
		}
	}

	// 第二次从后向前遍历，找到第一个比 k 小的数，将它们交换
	if idx >= 0 {
		for i := size - 1; i > idx; i-- {
			if nums[i] > nums[idx] {
				nums[idx], nums[i] = nums[i], nums[idx]
				break
			}
		}
	}

	// 从第一次遍历获得的数 k 的下一个坐标开始作为子数组的起点, 翻转子数组
	i := idx + 1
	j := size - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
