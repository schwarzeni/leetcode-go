package _1

// 2020.04.16
// https://leetcode-cn.com/problems/two-sum/
//func twoSum(nums []int, target int) []int {
//	copyArr := make([]int, len(nums))
//	copy(copyArr, nums)
//	sort.Ints(copyArr)
//	var (
//		i = 0
//		j = len(nums) - 1
//	)
//	for i < j {
//		r := copyArr[i] + copyArr[j]
//		if r > target {
//			j--
//		} else if r < target {
//			i++
//		} else {
//			ri := -1
//			rj := -1
//			for idx, v := range nums {
//				if v == copyArr[i] && ri == -1 {
//					ri = idx
//				} else if v == copyArr[j] && rj == -1 {
//					rj = idx
//				}
//			}
//			return []int{ri, rj}
//		}
//	}
//	return []int{0, len(nums) - 1}
//}
func twoSum(nums []int, target int) []int {
	var cache = make(map[int]int)

	for idx, v := range nums {
		r := target - v
		if i, ok := cache[r]; ok {
			return []int{i, idx}
		}
		cache[v] = idx
	}

	return []int{0, len(nums) - 1}
}
