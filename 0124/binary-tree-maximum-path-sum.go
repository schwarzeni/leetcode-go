package _0124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var getMax func(r *TreeNode) int
	var ans = math.MinInt64

	getMax = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := max(0, getMax(r.Left))
		right := max(0, getMax(r.Right))
		ans = max(ans, r.Val+left+right)
		return max(left, right) + r.Val
	}

	getMax(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
