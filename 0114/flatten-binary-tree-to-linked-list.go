package _0114

// 2020.06.13
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flatten2(root)
}

func flatten2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 左右都有孩子的情况
	if root.Left != nil && root.Right != nil {
		right := root.Right
		left := root.Left
		root.Left = nil
		root.Right = left
		flatten2(left).Right = right
		return flatten2(right)
	}
	// 只有右孩子的情况
	if root.Left == nil && root.Right != nil {
		return flatten2(root.Right)
	}
	// 只有左孩子的情况
	if root.Right == nil && root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		return flatten2(root.Right)
	}
	return root
}
