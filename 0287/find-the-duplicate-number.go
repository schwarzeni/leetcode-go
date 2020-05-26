package _0287

// 2020.05.26
// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 「Floyd 判圈算法」
// 先设置慢指针 slow 和快指针 fast ，慢指针每次走一步，快指针每次走两步
// 两个指针在有环的情况下一定会相遇
// 再将 slow 放置起点 0，两个指针每次同时移动一步，相遇的点就是答案
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
