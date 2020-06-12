package _0108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l, r := 0, len(nums)-1
	mid := l + (r-l)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[l:mid])
	root.Right = sortedArrayToBST(nums[mid+1 : r+1])
	return root
}
