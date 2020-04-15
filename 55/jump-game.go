package _55

// 2020.04.15
// https://leetcode-cn.com/problems/jump-game/
//func canJump(nums []int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//	cache := make([]int, len(nums))
//	cache[0] = 1
//
//	for idx, v := range nums {
//		if cache[idx] == 0 {
//			continue
//		}
//		for i := idx; i <= idx+v; i++ {
//			if i == len(nums)-1 {
//				return true
//			}
//			cache[i] = 1
//		}
//	}
//	return cache[len(nums)-1] == 1
//}

//      0 1 2 3 4
//     [2,3,1,1,4]
//
// 0    1 1 1 0 0 0
// 1      1 1 1 1 0
// 2        1 1 0 0
// 3          1 1 0
// 4            1 1

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	energy := 0

	for idx, v := range nums {
		energy -= 1
		if energy < v {
			energy = v
		}
		if idx == len(nums)-1 || energy+idx >= len(nums)-1 {
			return true
		}
		if energy <= 0 {
			return false
		}
	}
	return true
}
