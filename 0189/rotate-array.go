package _0189

// 2020.08.20
// https://leetcode-cn.com/problems/rotate-array/
// func rotate(nums []int, k int)  {
//     l := len(nums)
//     k = k % l
//     for i := 0; i < k; i++ {
//         v := nums[l-1]
//         for i := l-1 ; i >= 1; i-- {
//             nums[i] = nums[i-1]
//         }
//         nums[0] = v
//     }
// }

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
